package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBenin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm28 30H24v26.829c-.62-.185-1.232-.39-1.833-.616V5.787A27.86 27.86 0 0 1 32 4c15.439 0 28 12.561 28 28z"/>`),
		g.Group(children),
	)
}