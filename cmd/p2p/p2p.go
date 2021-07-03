package p2p

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"time"
	"os"
	"bufio"

	"github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	net "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	pstore "github.com/libp2p/go-libp2p-core/peerstore"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	ma "github.com/multiformats/go-multiaddr"

	mplex "github.com/libp2p/go-libp2p-mplex"
	direct "github.com/libp2p/go-libp2p-webrtc-direct"
	"github.com/pion/webrtc/v3"
)

// Message
type Message struct {
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

func (m Message) String() string {
	return m.Sender + ": " + m.Body
}

// Connection
type Connection struct {
	net.Stream
	buf * bufio.ReadWriter
	dec *json.Decoder
	enc *json.Encoder
}

func newConn(s net.Stream) *Connection {
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	return &Connection{
		Stream: s,
		buf: rw,
		dec:    json.NewDecoder(rw),
		enc:    json.NewEncoder(rw),
	}
}

func (conn *Connection) Read() (msg Message, err error) {
	err = conn.dec.Decode(&msg)
	return
}

func (conn *Connection) Write(msg Message) error {
	err := conn.enc.Encode(msg)
	conn.buf.Flush()
	return err
}

type Node struct {
	host.Host
	handlers map[protocol.ID]func(*Connection)
}

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It won't encrypt the connection if insecure is true.
func makeBasicHost(listenPort int, insecure bool, randseed int64) (host.Host, error) {

	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
	// deterministic randomness source to make generated keys stay the same
	// across multiple runs
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// Generate a key pair for this host. We will use it at least
	// to obtain a valid host ID.
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	transport := direct.NewTransport(
		webrtc.Configuration{},
		new(mplex.Transport),
	)

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/http/p2p-webrtc-direct", listenPort)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
		libp2p.Transport(transport),
	}

	if insecure {
		opts = append(opts, libp2p.NoSecurity)
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
	f, _ := os.Create("addr.txt")
	if insecure {
		fmt.Fprintf(f, "Now run \"./libp2p-echo -l %d -d %s -insecure\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		fmt.Fprintf(f, "Now run \"./libp2p-echo -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, nil
}

func NewNode(port int) (*Node, error) {
	h, err := makeBasicHost(port, true, time.Now().UnixNano())
	return &Node{
		Host:     h,
		handlers: make(map[protocol.ID]func(*Connection)),
	}, err
}

// Connect is used for simply connecting to another peer given its multiaddress.
func (n *Node) Connect(ctx context.Context, proto protocol.ID, addr string) (*Connection, error) {
	// The following code extracts target's the peer ID from the
	// given multiaddress
	ipfsaddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		log.Fatalln(err)
	}

	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Fatalln(err)
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		log.Fatalln(err)
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	n.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /echo/1.0.0 protocol
	s, err := n.NewStream(ctx, peerid, proto)
	return newConn(s), err
}

func (n *Node) Handle(path protocol.ID, h func(*Connection)) {
	n.handlers[path] = h
}

func (n *Node) Listen(ctx context.Context) error {
	// Set a stream handler on host A. path is
	// a user-defined protocol name.
	for path, handler := range n.handlers {
		n.SetStreamHandler(path, func(s net.Stream) {
			conn := newConn(s)
			handler(conn)
		})
	}

	select {} // hang forever
}
