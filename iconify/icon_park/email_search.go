package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 24V9H24H4V24V39H24"/><circle cx="36" cy="34" r="5" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 37L44 40"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 9L24 24L44 9"/></g>`),
		g.Group(children),
	)
}