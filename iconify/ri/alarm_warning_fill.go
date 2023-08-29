package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmWarningFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.001 20v-6a8 8 0 0 1 16 0v6h1v2h-18v-2h1Zm2-6h2a4 4 0 0 1 4-4V8a6 6 0 0 0-6 6Zm5-12h2v3h-2V2Zm8.778 2.808l1.414 1.414l-2.12 2.121l-1.415-1.414l2.121-2.121ZM2.81 6.222l1.414-1.414l2.121 2.12L4.93 8.344L2.809 6.222Z"/>`),
		g.Group(children),
	)
}