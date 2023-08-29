package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCameroon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM20 6.707a27.613 27.613 0 0 1 2-.854v52.294a28.53 28.53 0 0 1-2-.854V6.707zM37.254 40L32 36.246L26.746 40l1.984-6.105l-5.23-3.785l6.479-.018L32 24l2.021 6.092l6.479.018l-5.232 3.785L37.254 40zM42 58.147V5.853C52.514 9.888 60 20.083 60 32s-7.486 22.112-18 26.147z"/>`),
		g.Group(children),
	)
}