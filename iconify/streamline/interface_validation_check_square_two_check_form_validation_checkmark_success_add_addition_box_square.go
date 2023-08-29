package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceValidationCheckSquareTwoCheckFormValidationCheckmarkSuccessAddAdditionBoxSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="m3 7l1.5 1.5L7 5m1 2.5h3"/></g>`),
		g.Group(children),
	)
}