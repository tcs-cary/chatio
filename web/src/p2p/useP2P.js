import Libp2p from "libp2p";
import Websockets from "libp2p-websockets";
import WebRTCStar from "libp2p-webrtc-star";
import { NOISE } from "libp2p-noise";
import Mplex from "libp2p-mplex";
import Bootstrap from "libp2p-bootstrap";
import { ref, watchEffect, unref } from "vue";
import pipe from "it-pipe";
import lp from "it-length-prefixed";
// const multiaddr = require('multiaddr');

export async function useConnect() {
  // Create our libp2p node

  const libp2p = await Libp2p.create({
    addresses: {
      // Add the signaling server address, along with our PeerId to our multiaddrs list
      // libp2p will automatically attempt to dial to the signaling server so that it can
      // receive inbound connections from other peers
      listen: [
        "/ip4/127.0.0.1/tcp/9090/wss/p2p-webrtc-star",
      ]
    },
    modules: {
      transport: [Websockets, WebRTCStar],
      connEncryption: [NOISE],
      streamMuxer: [Mplex],
      peerDiscovery: [Bootstrap]
    },
    config: {
      peerDiscovery: {
        // The `tag` property will be searched when creating the instance of your Peer Discovery service.
        // The associated object, will be passed to the service when it is instantiated.
        [Bootstrap.tag]: {
          enabled: false,
          list: [
            "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
            "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
            "/dnsaddr/bootstrap.libp2p.io/p2p/QmZa1sAxajnQjVM8WjWXoMbmPd7NsWhfKsPkErzpm9wGkp",
            "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
            "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt"
          ]
        }
      }
    }
  })



  console.log("created");
  // Listen for new peers
  libp2p.on("peer:discovery", peerId => {
    console.log(`Found peer ${peerId.toB58String()}`);
  });

  // Listen for new connections to peers
  libp2p.connectionManager.on("peer:connect", async connection => {


    console.log(`Connected to ${connection.remotePeer.toB58String()}`);

    const addresses = libp2p.peerStore.addressBook.get(connection.remotePeer);

    if (addresses === undefined) return;

    const { stream } = await libp2p.dialProtocol(connection.remotePeer, "/chatio/1.0.0");

    //pipe(outgoingMessages, lp.encode(), stream.sink);





    // const ints = (async function*() {
    //   let i = 0
    //   while (true) yield i++
    // })

    console.log("dialed");

    watchEffect(async () => {
      const msg = unref(outgoingMessages.value[outgoingMessages.value.length - 1]);
      console.log(msg);
      await pipe([JSON.stringify(msg)], lp.encode(), stream);
    })

    console.log("sent");
  });



  // Listen for peers disconnecting
  libp2p.connectionManager.on("peer:disconnect", connection => {
    console.log(`Disconnected from ${connection.remotePeer.toB58String()}`);
  });

  let incomingMessages = ref([]);

  let outgoingMessages = ref([]);

  await libp2p.handle('/chatio/1.0.0', async ({ stream }) => {
    // Send stdin to the stream
    console.log("handle");
    console.log(stream);


    // pipe(outgoingMessages.value, lp.encode(), stream.sink);
    //
    pipe(
      // Read from the stream (the source)
      stream.source,
      // Decode length-prefixed data
      lp.decode(),
      // Sink function
      async function(source) {
        // For each chunk of data
        for await (const msg of source) {
          // Output the data as a utf8 string
          incomingMessages.value.push(JSON.parse(msg.toString()));
          console.log("recived     " + msg.toString());

        }
      }
    )

  })

  await libp2p.start();
  console.log(`libp2p id is ${libp2p.peerId.toB58String()}`);

  // const addr = multiaddr('/dns4/wrtc-star1.par.dwebops.pub/tcp/443/wss/p2p-webrtc-star/p2p/QmTyDvctiAzHif8av2X23SoSU249KngGBw7Tp6vVJ5di4H');


  return {
    onMessage: f => watchEffect(() => {
      incomingMessages.value.forEach(message => {
        console.log(message);
      })
      if (incomingMessages.value.length == 0) return;
      const msg = unref(incomingMessages.value[incomingMessages.value.length - 1]);
      f(msg);
    }),
    send: msg => outgoingMessages.value.push(msg)
  };
}