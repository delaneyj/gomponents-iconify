package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rega(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.384 17.5H13.181l3.384-13h14.202l-3.383 13zm10.818 13H24l3.384-13h14.202l-3.384 13zm-17.586 13H6.414l3.384-13H24l-3.384 13z"/>`),
		g.Group(children),
	)
}