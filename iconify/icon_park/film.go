package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path stroke-linecap="round" d="M16 6V42"/><path stroke-linecap="round" d="M32 6V42"/><path stroke-linecap="round" d="M6 15H16"/><path stroke-linecap="round" d="M32 15H42"/><path stroke-linecap="round" d="M6 33H16"/><path stroke-linecap="round" d="M6 24H42"/><path stroke-linecap="round" d="M32 33H42"/></g>`),
		g.Group(children),
	)
}