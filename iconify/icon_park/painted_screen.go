package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintedScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="28" x="5" y="10" fill="#2F88FF" stroke="#000" stroke-width="4" rx="3"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 24L13 17"/><rect width="4" height="4" x="15" y="29" fill="#fff" rx="2" transform="rotate(90 15 29)"/></g>`),
		g.Group(children),
	)
}