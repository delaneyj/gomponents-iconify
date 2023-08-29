package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerLiquidThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.125 20.076a6.505 6.505 0 0 1-8-6.326c0-4.875 6.5-13 6.5-13a46.389 46.389 0 0 1 5.306 8.5"/><path d="M22.875 15.75a7.669 7.669 0 0 1-6 7.5a7.669 7.669 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5v2.25Zm-6-.75v4.5m-2.25-2.25h4.5m-11.5 0a3.5 3.5 0 0 1-3.5-3.5"/></g>`),
		g.Group(children),
	)
}