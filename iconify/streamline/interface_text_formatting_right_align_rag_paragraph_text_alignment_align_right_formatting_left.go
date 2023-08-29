package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingRightAlignRagParagraphTextAlignmentAlignRightFormattingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 3h-10m10 3h-7m7 6H.5m13-3h-10"/>`),
		g.Group(children),
	)
}