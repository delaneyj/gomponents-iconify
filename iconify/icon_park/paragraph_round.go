package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 9H42"/><path d="M24 19H42"/><path d="M6 29H42"/><path d="M6 39H42"/><circle cx="12" cy="14" r="5" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}