package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#2F88FF" d="M24 4C26.6264 4 29.2272 4.51732 31.6537 5.52241C34.0802 6.5275 36.285 8.00069 38.1421 9.85787C39.9993 11.715 41.4725 13.9198 42.4776 16.3463C43.4827 18.7728 44 21.3736 44 24L24 24V4Z"/></g>`),
		g.Group(children),
	)
}