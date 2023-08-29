package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingJustifiedAlignJustifiedAlignAlignmentParagraphFormattingText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 1H.5m13 3H.5m13 3H.5m13 6H.5m13-3H.5"/>`),
		g.Group(children),
	)
}