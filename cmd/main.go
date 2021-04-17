package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"chatio/p2p"
	"chatio/ui"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/textinput"
)

func main() {
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	n, err := p2p.NewNode(*listenF)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := n.Connect(ctx, "/echo/1.0.0", *target)
	if err != nil {
		fmt.Println(err)
		return
	}

	n.Handle("/echo/1.0.0", func(c *p2p.Connection) {
		msg, err := c.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(msg)
		c.Write(msg)
	})

	if *target == "" {
		if err = n.Listen(ctx); err != nil {
			fmt.Println(err)
		}
		return
	}

	err = conn.Write(p2p.Message{Sender: "Me", Body: "Hello, World!"})
	if err != nil {
		fmt.Println(err)
		return
	}

	msg, err := conn.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("received message:", msg)

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
}
