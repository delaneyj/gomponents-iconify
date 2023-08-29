package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfferentFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21 5H10a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V24.75"/><path d="M33 24H21V12m0 12L39 6"/></g>`),
		g.Group(children),
	)
}