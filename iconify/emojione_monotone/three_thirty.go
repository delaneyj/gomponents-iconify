package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m16.973 36.5l-14.061-3.772a4 4 0 0 1-.912.718V58h-4V35.445a3.988 3.988 0 0 1-1.889-2.543L26 32.336l1.023-3.836l2.235.6c.224-.211.474-.389.741-.545V26h4v2.555a3.973 3.973 0 0 1 1.814 2.305L50 34.666L48.973 38.5"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}