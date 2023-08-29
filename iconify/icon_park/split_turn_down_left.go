package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitTurnDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M37 21.9999H21C16.5817 21.9999 13 25.5817 13 29.9999V43.9999"/><circle cx="37" cy="8.944" r="5" fill="#2F88FF" transform="rotate(-90 37 8.944)"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 13.9999V42.9999"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 38.9999L37 43.9999L32 38.9999"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 38.9999L13 43.9999L8 38.9999"/></g>`),
		g.Group(children),
	)
}