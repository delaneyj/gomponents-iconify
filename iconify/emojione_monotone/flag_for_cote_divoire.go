package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCoteDivoire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-9.834 56.213V5.787C25.228 4.635 28.541 4 32 4s6.771.635 9.833 1.787v52.426C38.771 59.365 35.459 60 32 60s-6.772-.635-9.834-1.787z"/>`),
		g.Group(children),
	)
}