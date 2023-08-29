package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingIndentRightAlignmentAlignIndentParagraphFormattingRightText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1.5h6m-4 4h4m-6 4h6m-4 4h4m7-8.5l-2 2l2 2M9 13.5V.5"/>`),
		g.Group(children),
	)
}