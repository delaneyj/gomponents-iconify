package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingIndentIncreaseAlignmentAlignIndentParagraphIncreaseFormattingText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 4H6m7.5 4H6m7.5 4H.5M2 5l2 2l-2 2"/>`),
		g.Group(children),
	)
}