package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4 30c0 1.477-.81 2.752-2 3.445V58h-4V36.857L28.857 38L26 35.145l2.142-2.143A3.993 3.993 0 0 1 28 32c0-1.477.81-2.752 2-3.445V26h4v1.144L43.143 18L46 20.859L35.858 31c.084.321.142.652.142 1z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}