package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugsurfing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.712 4.344L13.5 25.674l11.06-.262l-7.346 17.932L34.5 20.657l-11.883.294Z"/>`),
		g.Group(children),
	)
}