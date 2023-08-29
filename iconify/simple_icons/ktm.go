package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ktm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 15.735h3.354l.843-2.06l1.55 2.06h7.225l2.234-2.081l-.372 2.081h2.83L20 13.675l-.32 2.06h3.052L24 9.99h-3.068l-2.486 2.191l.48-2.19h-2.942l-3.209 3.216l1.342-3.938h4.907l.225-1.003H6.381l-.378 1.003h4.732l-1.994 5.054l-1.572-2.066L9.886 9.99H7.612l-2.787 2.23l.938-2.23H2.44L0 15.735Z"/>`),
		g.Group(children),
	)
}