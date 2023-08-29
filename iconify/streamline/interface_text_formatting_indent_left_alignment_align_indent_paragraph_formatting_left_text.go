package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingIndentLeftAlignmentAlignIndentParagraphFormattingLeftText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1.5h-6m4 4h-4m6 4h-6m4 4h-4M.5 5l2 2l-2 2M5 13.5V.5"/>`),
		g.Group(children),
	)
}