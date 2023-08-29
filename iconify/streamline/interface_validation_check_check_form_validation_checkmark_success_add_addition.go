package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceValidationCheckCheckFormValidationCheckmarkSuccessAddAddition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 8.55l2.73 3.51a1 1 0 0 0 .78.39a1 1 0 0 0 .78-.36L13.5 1.55"/>`),
		g.Group(children),
	)
}