package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForYemen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM5.787 22.167h52.426C59.365 25.229 60 28.541 60 32s-.635 6.771-1.787 9.833H5.787C4.635 38.771 4 35.459 4 32s.635-6.771 1.787-9.833z"/>`),
		g.Group(children),
	)
}