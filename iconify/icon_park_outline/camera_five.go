package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="21" r="16"/><circle cx="24" cy="21" r="7"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 43h16m-8-6v6"/></g>`),
		g.Group(children),
	)
}