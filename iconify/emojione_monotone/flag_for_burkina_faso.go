package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBurkinaFaso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-8.85 27.91l6.744-.021L32 23.15l2.105 6.739l6.744.021l-5.448 4.186l2.068 6.754L32 36.695l-5.47 4.154l2.065-6.754l-5.445-4.185zM32 60C16.561 60 4 47.439 4 32h19.408l3.449 2.65l-1.762 5.762l-1.358 4.443l3.7-2.811L32 38.579l4.563 3.466l3.703 2.813l-1.361-4.446l-1.764-5.761l3.45-2.65H60C60 47.439 47.439 60 32 60z"/>`),
		g.Group(children),
	)
}