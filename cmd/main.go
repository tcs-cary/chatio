package main

import (
	"context"
	"fmt"
	"time"

	"chatio/ui"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	// "github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"

	"github.com/pion/webrtc/v3"
	"chatio/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	updateText := make(chan string, 1)

	router := ui.NewRouter()

	login, err := loginView(router, cancel, updateText)
	if err != nil {
		panic(err)
	}

	chat, err := chatRoomView(updateText)
	if err != nil {
		panic(err)
	}

	router.AddRoute("/", login)
	router.AddRoute("/chat", chat)

	userInterface := ui.New(router, 16*time.Millisecond) // 16 ms ~ 60 fps

	err = userInterface.Run(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Good Bye!")
}

func chatRoomView(updateText <-chan string) (*ui.View, error) {
	input, err := textinput.New(
		textinput.LabelAlign(align.HorizontalLeft),
		textinput.PlaceHolder("Enter your message. Now."),
		textinput.FillColor(cell.ColorBlack),
		textinput.ClearOnSubmit(),
		textinput.OnSubmit(func(message string) error {
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}
	// go func() {
	// 	if err := unicode.Write(<-updateText); err != nil {
	// 		panic(err)
	// 	}
	// }()

	builder := grid.New()
	sendButton, err := button.New("Send", func() error {
		return nil
	})
	builder.Add(
		grid.RowHeightPercWithOpts(80, []container.Option{container.Border(linestyle.Light)}),
		grid.RowHeightPerc(20,
			grid.ColWidthPercWithOpts(90, []container.Option{container.Border(linestyle.Light)}, grid.Widget(
				input,
			)),
			grid.ColWidthPercWithOpts(10, []container.Option{container.Border(linestyle.Light)}, grid.Widget(sendButton)),
		),
	)
	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}

	chatView := ui.NewView(gridOpts...)
	return chatView, nil
}

func loginView(router *ui.Router, cancel func(), updateText chan<- string) (*ui.View, error) {
	input, err := textinput.New(
		textinput.Label("New text:", cell.FgColor(cell.ColorNumber(33))),
		textinput.MaxWidthCells(20),
		textinput.Border(linestyle.Light),
		textinput.PlaceHolder("Enter any text"),
	)
	if err != nil {
		return nil, err
	}

	submitB, err := button.New("Submit", func() error {
		updateText <- input.ReadAndClear()
		close(updateText)
		router.NavigateTo("/chat")
		return nil
	},
		button.GlobalKey(keyboard.KeyEnter),
		button.FillColor(cell.ColorNumber(220)),
	)
	if err != nil {
		return nil, err
	}
	clearB, err := button.New("Clear", func() error {
		input.ReadAndClear()
		return nil
	},
		button.WidthFor("Submit"),
		button.FillColor(cell.ColorNumber(220)),
	)
	if err != nil {
		return nil, err
	}
	quitB, err := button.New("Quit", func() error {
		cancel()
		close(updateText)
		return nil
	},
		button.WidthFor("Submit"),
		button.FillColor(cell.ColorNumber(196)),
	)
	if err != nil {
		return nil, err
	}

	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(20,
			grid.Widget(
				input,
				container.AlignHorizontal(align.HorizontalCenter),
				container.AlignVertical(align.VerticalBottom),
				container.MarginBottom(1),
			),
		),
	)

	builder.Add(
		grid.RowHeightPerc(40,
			grid.ColWidthPerc(20),
			grid.ColWidthPerc(20,
				grid.Widget(
					submitB,
					container.AlignVertical(align.VerticalTop),
					container.AlignHorizontal(align.HorizontalRight),
				),
			),
			grid.ColWidthPerc(20,
				grid.Widget(
					clearB,
					container.AlignVertical(align.VerticalTop),
					container.AlignHorizontal(align.HorizontalCenter),
				),
			),
			grid.ColWidthPerc(20,
				grid.Widget(
					quitB,
					container.AlignVertical(align.VerticalTop),
					container.AlignHorizontal(align.HorizontalLeft),
				),
			),
			grid.ColWidthPerc(20),
		),
	)

	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}

	loginView := ui.NewView(gridOpts...)
	return loginView, nil

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", d.Label(), d.ID())

			for range time.NewTicker(5 * time.Second).C {
				message := signal.RandSeq(15)
				fmt.Printf("Sending '%s'\n", message)

				// Send the message as text
				sendErr := d.SendText(message)
				if sendErr != nil {
					panic(sendErr)
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// Wait for the offer to be pasted
	offer := webrtc.SessionDescription{}
	signal.Decode(signal.MustReadStdin(), &offer)

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	// Create an answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	// Output the answer in base64 so we can paste it in browser
	fmt.Println(signal.Encode(*peerConnection.LocalDescription()))

	// Block forever
	select {}
}
