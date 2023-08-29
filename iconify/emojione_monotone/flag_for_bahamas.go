package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBahamas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-8.419 39.833L12.934 52.479a28.144 28.144 0 0 1-1.414-1.414L30.586 32L11.52 12.934a28.57 28.57 0 0 1 1.414-1.414L23.58 22.167h34.633C59.365 25.228 60 28.54 60 32s-.635 6.772-1.787 9.833H23.581z"/>`),
		g.Group(children),
	)
}