package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDeleteSquareButtonRemoveButtonsAddSquareDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.12 4.88L4.88 9.12m0-4.24l4.24 4.24"/><rect width="13" height="13" x=".5" y=".5" rx="3"/></g>`),
		g.Group(children),
	)
}