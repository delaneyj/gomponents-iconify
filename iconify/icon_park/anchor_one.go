package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M40 35C40 25.7953 32.8366 10 24 10C15.1634 10 8 25.7953 8 35"/><rect width="8" height="8" x="4" y="35" fill="#2F88FF"/><rect width="8" height="8" x="4" y="6" fill="#2F88FF"/><rect width="8" height="8" x="36" y="35" fill="#2F88FF"/><rect width="8" height="8" x="36" y="6" fill="#2F88FF"/><path stroke-linecap="round" d="M12 10H36"/></g>`),
		g.Group(children),
	)
}