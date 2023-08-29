package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path d="M6 14H40"/><path d="M30 22H42"/><path d="M30 30H42"/><path d="M23 22H24"/><path d="M23 30H24"/><path d="M14 6L14 42"/></g>`),
		g.Group(children),
	)
}