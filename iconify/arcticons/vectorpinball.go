package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vectorpinball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.44 35.61l22.06-8.86V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v17.74l11.67-4.69M5.5 25.24v15.31a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2v-13.8M27.78 11.1a2.51 2.51 0 1 1-2.51 2.51a2.51 2.51 0 0 1 2.51-2.51Z"/>`),
		g.Group(children),
	)
}