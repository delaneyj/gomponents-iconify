package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingImageTopInsertParagraphImageTextAlignTopFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 13.5h-9m13-3H.5m13-3H.5"/><rect width="4" height="13" x="5" y="-4" rx=".5" transform="rotate(-90 7 2.5)"/></g>`),
		g.Group(children),
	)
}