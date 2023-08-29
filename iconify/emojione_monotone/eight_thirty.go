package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m2 33.445V58h-4V35.445a4.036 4.036 0 0 1-.913-.719L15.025 38.5L14 34.664l14.186-3.807A3.983 3.983 0 0 1 30 28.554V26h4v2.554c.268.156.519.334.741.545l2.231-.599L38 32.334l-2.11.566A3.988 3.988 0 0 1 34 35.445"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}