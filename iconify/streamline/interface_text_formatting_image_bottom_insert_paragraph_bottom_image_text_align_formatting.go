package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingImageBottomInsertParagraphBottomImageTextAlignFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 6.5h-9m13-3H.5m13-3H.5"/><rect width="4" height="13" x="5" y="5" rx=".5" transform="rotate(-90 7 11.5)"/></g>`),
		g.Group(children),
	)
}