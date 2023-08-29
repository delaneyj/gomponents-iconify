package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGreenland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M24.133 17.25c-8.145 0-14.75 6.604-14.75 14.75h29.5c0-8.146-6.605-14.75-14.75-14.75"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 2c15.439 0 28 12.561 28 28H38.883c0 8.145-6.604 14.75-14.75 14.75S9.383 40.145 9.383 32H4C4 16.561 16.561 4 32 4z"/>`),
		g.Group(children),
	)
}