package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vectorifydahome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m31.616 13.799l2.585 2.585a1.02 1.02 0 0 1-.012 1.455l-2.467 2.373l-3.934-3.934l2.373-2.467a1.02 1.02 0 0 1 1.455-.012Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M13.5 34.5v-3.934l14.288-14.288l3.934 3.934L17.434 34.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}