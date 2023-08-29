package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphBreakTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 9h36M6 19h36M6 29h18M6 39h10m8 0h13a5 5 0 0 0 0-10h-5m-8 10l4-4m-4 4l4 4"/>`),
		g.Group(children),
	)
}