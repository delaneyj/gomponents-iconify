package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 9H42"/><path d="M24 19H42"/><path d="M6 29H42"/><path d="M6 39H42"/><rect width="10" height="10" x="6" y="9" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}