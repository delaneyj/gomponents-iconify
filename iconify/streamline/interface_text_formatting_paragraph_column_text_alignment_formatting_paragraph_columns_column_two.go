package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingParagraphColumnTextAlignmentFormattingParagraphColumnsColumnTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 1H2m3.5 3h-5m5 3h-5m5 6h-5m5-3h-5m13-9h-5m5 3h-5m5 3h-5m3.5 6H8.5m5-3h-5"/>`),
		g.Group(children),
	)
}