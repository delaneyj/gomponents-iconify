package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFaroeIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 2c12.683 0 23.391 8.436 26.834 20H32V4zm-4 .289a27.973 27.973 0 0 1 2-.21V26h29.349c.144.659.266 1.325.362 2H28V4.289zM16 9.027V24H5.166A28.045 28.045 0 0 1 16 9.027zm0 45.946A28.045 28.045 0 0 1 5.166 40H16v14.973zm4 2.329a27.608 27.608 0 0 1-2-1.057V38H4.651a27.407 27.407 0 0 1-.362-2H20v21.302zM20 28H4.289c.097-.675.218-1.341.362-2H18V7.755a28.008 28.008 0 0 1 2-1.057V28zm12 32V40h26.834C55.391 51.564 44.683 60 32 60zm27.349-22H30v21.921a27.973 27.973 0 0 1-2-.21V36h31.711a28.523 28.523 0 0 1-.362 2z"/>`),
		g.Group(children),
	)
}