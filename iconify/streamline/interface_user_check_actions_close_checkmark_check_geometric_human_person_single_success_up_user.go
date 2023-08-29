package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceUserCheckActionsCloseCheckmarkCheckGeometricHumanPersonSingleSuccessUpUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 8l-4.12 5.5l-2.75-2.06"/><circle cx="5" cy="2.75" r="2.25"/><path d="M3 12.5H.5V11a4.5 4.5 0 0 1 7.68-3.18"/></g>`),
		g.Group(children),
	)
}