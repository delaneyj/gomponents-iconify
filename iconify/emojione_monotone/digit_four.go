package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M33.662 36.217V25.326l-7.674 10.891z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm12 39.586h-4.145V48h-6.193v-6.414H20v-5.348L34.48 16h5.375v20.217H44v5.369z"/>`),
		g.Group(children),
	)
}