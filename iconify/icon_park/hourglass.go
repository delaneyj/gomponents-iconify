package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 4H40"/><path d="M8 44H40"/><path d="M12 4V16L21 26"/><path d="M36 44V29.5L27 21"/><path d="M12 44V30L18.5 23.5"/><path d="M36 4V16L29.5 23.5"/><path d="M18 33H19"/><path d="M29.1465 32.6465L29.8536 33.3536"/><path d="M24 38H25"/></g>`),
		g.Group(children),
	)
}