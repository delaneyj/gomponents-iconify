package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailDelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H24H4V24V39H24"/><path d="M4 9L24 24L44 9"/><path fill="#2F88FF" d="M32 31H42L40 41H34L32 31Z"/><path d="M36 31L38 27"/></g>`),
		g.Group(children),
	)
}