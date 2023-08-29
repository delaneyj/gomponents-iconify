package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceRemoveSquareSubtractButtonsRemoveAddButtonSquareDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 7h6"/><rect width="13" height="13" x=".5" y=".5" rx="3"/></g>`),
		g.Group(children),
	)
}