package ui

import (
	"context"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// TODO(Zaba505): Is this useful?
type Widget interface {
	widgetapi.Widget

	Listen(event string, handler func(data interface{}) error)

	Bind(name string, data interface{})
}

type View struct {
	opts []container.Option
}

func NewView(opts ...container.Option) *View {
	return &View{
		opts: opts,
	}
}

type Router struct {
	term   terminalapi.Terminal
	routes map[string]*View

	pathMu  sync.Mutex
	curPath string
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]*View),
	}
}

func (r *Router) CurrentView() *View {
	r.pathMu.Lock()
	defer r.pathMu.Unlock()
	return r.routes[r.curPath]
}

func (r *Router) AddRoute(path string, view *View) {
	r.routes[path] = view
}

func (r *Router) NavigateTo(path string) {
	r.pathMu.Lock()
	defer r.pathMu.Unlock()

	r.curPath = path
}

type UI struct {
	router         *Router
	redrawInterval time.Duration
}

func New(router *Router, redrawInterval time.Duration) UI {
	return UI{
		router,
		redrawInterval,
	}
}

const rootId = "root"

func (u UI) Run(ctx context.Context) error {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ticker := time.NewTicker(u.redrawInterval)
	defer ticker.Stop()

	u.router.term = t
	u.router.NavigateTo("/")

	view := u.router.CurrentView()

	opts := make([]container.Option, len(view.opts)+1)
	copy(opts, view.opts)
	opts[len(opts)-1] = container.ID(rootId)

	cont, err := container.New(t, opts...)
	if err != nil {
		return err
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}

	ctrl, err := termdash.NewController(t, cont, termdash.KeyboardSubscriber(quitter))
	if err != nil {
		return err
	}

	curPath := u.router.curPath
	for {
		view := u.router.CurrentView()

		u.router.pathMu.Lock()
		p := u.router.curPath
		u.router.pathMu.Unlock()

		if curPath != p {
			curPath = p
			cont.Update(rootId, view.opts...)
		}

		select {
		case <-cctx.Done():
			ctrl.Close()
			return nil
		case <-ticker.C:
			if err := ctrl.Redraw(); err != nil {
				return err
			}
		}
	}
}
