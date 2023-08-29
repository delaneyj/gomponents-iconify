package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwelveThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4 30c0 1.477-.81 2.753-2 3.445V58h-4V38.51l-2-.533l.925-3.447A3.962 3.962 0 0 1 28 32c0-1.479.81-2.753 2-3.445V26h1.213l2.951-11L38 16.027L34.538 28.93A3.976 3.976 0 0 1 36 32z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}