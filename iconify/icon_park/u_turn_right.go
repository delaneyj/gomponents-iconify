package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M34 13H15C8.92487 13 4 17.9249 4 24V24C4 30.0751 8.92487 35 15 35H44"/><path stroke-linecap="round" stroke-linejoin="round" d="M39 30L44 35L39 40"/><circle cx="39" cy="13" r="5" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}