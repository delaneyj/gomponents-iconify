package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculatorpluspluswindowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.15 37.41a13.41 13.41 0 1 1 9.47-22.88l-4.52 4.52a7 7 0 1 0 0 9.91h0l4.52 4.52a13.37 13.37 0 0 1-9.47 3.93Zm3.23-16.14v5m-2.5-2.5h5m4.89-2.5v5m-2.5-2.5h5m3.04-16.08l-2.4 2.4m0-2.4l2.4 2.4M35.72 5.5v6.79h6.78"/>`),
		g.Group(children),
	)
}