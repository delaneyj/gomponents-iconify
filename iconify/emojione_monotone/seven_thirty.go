package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4 30c0 1.477-.81 2.753-2 3.445V58h-4V36.855L20.857 46L18 43.141L28.142 33A3.926 3.926 0 0 1 28 32c0-1.478.81-2.753 2-3.446V26h4v1.143L35.143 26L38 28.855l-2.142 2.142c.083.323.142.655.142 1.003z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}