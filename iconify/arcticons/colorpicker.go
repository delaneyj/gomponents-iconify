package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Colorpicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V8.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 8c2.5-.66 3-2.87 5-5s5.46 1.51 3.5 3.5s-4.35 2.55-4.9 4.9m-2 2L25 26.1a2.7 2.7 0 0 1-2.6.4s-1.25 1.38-2 .5s.5-1.8.5-1.8a2.78 2.78 0 0 1 .5-2.6L34 10m-1 8h-7m12.8-3.6l1.9-1.9L34.8 7l-1.9 1.9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.1 26l2.14 2.22a.7.7 0 0 0 1 0l1.38-1.38a.72.72 0 0 1 1 0l7.81 7.81a.7.7 0 0 1-.5 1.2H10.1a.7.7 0 0 1-.58-1.1L19 21.13a.7.7 0 0 1 1-.13l1.48 1.5"/>`),
		g.Group(children),
	)
}