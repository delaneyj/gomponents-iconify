package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mlmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 5.5v17.2l1.9-1.8l2.5 2.4l2.5-2.5l2.5 2.6l2.5-2.5l2.5 2.4l2.5-2.5l2 2V5.5m1 30.1c0-3.3 5.2-3.5 5.2 0m-5.2 0h5.2v4.3h-5.2zm3.8 4.4v2.5M35.7 40v2.5m-2.8-7.3v4m8.2-4v4m-5-6L35 31.7m2.9 1.5l1.1-1.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.6 9.7v8h4m-14.3 0v-8l4 8l4-8v8M37 28.6a8.5 8.5 0 1 0 8.5 8.5h0a8.49 8.49 0 0 0-8.5-8.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.5 42.5h-23a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v23.1"/>`),
		g.Group(children),
	)
}