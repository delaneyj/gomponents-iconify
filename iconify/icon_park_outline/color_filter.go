package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 40.944A11.955 11.955 0 0 0 32 44c6.627 0 12-5.373 12-12c0-5.591-3.824-10.29-9-11.622"/><path d="M13 20.378C7.824 21.71 4 26.408 4 32c0 6.627 5.373 12 12 12s12-5.373 12-12c0-1.55-.294-3.03-.828-4.39"/><path d="M24 28c6.627 0 12-5.373 12-12S30.627 4 24 4S12 9.373 12 16s5.373 12 12 12Z"/></g>`),
		g.Group(children),
	)
}