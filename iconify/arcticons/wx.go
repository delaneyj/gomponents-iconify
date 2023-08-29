package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.68 16.48l3.77 5.62l-3.77 5.62m7.53-11.24l-3.76 5.62l3.76 5.62m-20.42-7.44l2.8 7.44m2.81-7.44l-2.81 7.44m2.81-7.44l2.8 7.44m2.8-7.44l-2.8 7.44"/>`),
		g.Group(children),
	)
}