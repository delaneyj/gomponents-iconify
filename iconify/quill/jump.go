package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M26.622 5.17c.096-.043.185.072.119.154l-7.355 9.193a.5.5 0 0 0 .167.76l3.563 1.781a1 1 0 0 1-.037 1.806L5.38 26.83c-.096.043-.185-.072-.12-.154l7.355-9.193a.5.5 0 0 0-.167-.76l-3.563-1.781a1 1 0 0 1 .037-1.806l17.7-7.966Z"/>`),
		g.Group(children),
	)
}