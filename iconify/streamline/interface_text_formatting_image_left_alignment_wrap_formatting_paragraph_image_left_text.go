package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingImageLeftAlignmentWrapFormattingParagraphImageLeftText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 1H9m4.5 3H9m4.5 3H9m4.5 6H.5m13-3H.5"/><rect width="6" height="6" x=".5" y="1" rx=".5"/></g>`),
		g.Group(children),
	)
}