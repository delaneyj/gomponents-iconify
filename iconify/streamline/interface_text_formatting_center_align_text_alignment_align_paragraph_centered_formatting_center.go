package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingCenterAlignTextAlignmentAlignParagraphCenteredFormattingCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 1h13M2 4h10M3.5 7h7m-10 6h13M2 10h10"/>`),
		g.Group(children),
	)
}