package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerGapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24ZM48 136a8 8 0 0 1 0-16h24a8 8 0 0 1 0 16Zm46.06 37.25l-17 17a8 8 0 0 1-11.32-11.32l17-17a8 8 0 0 1 11.31 11.31Zm0-79.19a8 8 0 0 1-11.31 0l-17-17a8 8 0 0 1 11.34-11.29l17 17a8 8 0 0 1-.03 11.29ZM136 208a8 8 0 0 1-16 0v-24a8 8 0 0 1 16 0Zm0-136a8 8 0 0 1-16 0V48a8 8 0 0 1 16 0Zm54.23 118.23a8 8 0 0 1-11.32 0l-17-17a8 8 0 0 1 11.31-11.31l17 17a8 8 0 0 1 .01 11.31ZM208 136h-24a8 8 0 0 1 0-16h24a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}