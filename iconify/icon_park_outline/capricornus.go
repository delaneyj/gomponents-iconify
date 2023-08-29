package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capricornus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M18 11a6 6 0 0 0-12 0m12 0v18m12-18a6 6 0 0 0-12 0m12 0v24.75S30 43 22 43"/><circle cx="36" cy="30" r="6"/></g>`),
		g.Group(children),
	)
}