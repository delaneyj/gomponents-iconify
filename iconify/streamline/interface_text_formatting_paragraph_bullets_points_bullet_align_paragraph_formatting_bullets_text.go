package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingParagraphBulletsPointsBulletAlignParagraphFormattingBulletsText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 1H8m5.5 3H8m5.5 3H8m5.5 6H8m5.5-3H8"/><rect width="5" height="5" x=".5" y="1" rx=".5" transform="rotate(-90 3 3.5)"/><rect width="5" height="5" x=".5" y="8" rx=".5" transform="rotate(-90 3 10.5)"/></g>`),
		g.Group(children),
	)
}