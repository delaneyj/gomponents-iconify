package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBelgium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-9.833 56.213V5.787A27.86 27.86 0 0 1 32 4c3.459 0 6.771.635 9.834 1.787v52.426A27.873 27.873 0 0 1 32 60a27.86 27.86 0 0 1-9.833-1.787z"/>`),
		g.Group(children),
	)
}