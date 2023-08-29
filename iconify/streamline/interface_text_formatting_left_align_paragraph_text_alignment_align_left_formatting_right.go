package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingLeftAlignParagraphTextAlignmentAlignLeftFormattingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1h13M.5 4h10M.5 7h7m-7 6h13m-13-3h10"/>`),
		g.Group(children),
	)
}