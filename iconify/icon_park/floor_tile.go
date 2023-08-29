package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorTile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path stroke-linecap="round" d="M28 6L6 28"/><path stroke-linecap="round" d="M42 20L20 42"/><path stroke-linecap="round" d="M40 8L8 40"/><path stroke-linecap="round" d="M12 22L19 29"/><path stroke-linecap="round" d="M29 19L36 26"/></g>`),
		g.Group(children),
	)
}