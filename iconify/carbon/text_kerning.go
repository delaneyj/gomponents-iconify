package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextKerning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 24H18.83l2.29-2.29l-1.41-1.42L15 25l4.71 4.71l1.41-1.42L18.83 26H30v-2zm-16-3l6-17h-2l-6 17h2zM13 4L9 16L5 4H3l5 14h2l5-14h-2zm15 14h2L25 4h-2l-5 14h2l1-3h6zm-6.33-5L24 6l2.33 7z"/>`),
		g.Group(children),
	)
}