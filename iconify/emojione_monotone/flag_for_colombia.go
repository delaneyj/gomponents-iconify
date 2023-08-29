package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForColombia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m0 2c15.439 0 28 12.561 28 28H4C4 16.561 16.561 4 32 4m25 41H7c-.336-.654-.715-1.318-1-2h52c-.285.682-.664 1.346-1 2"/>`),
		g.Group(children),
	)
}