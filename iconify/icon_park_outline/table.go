package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" rx="2"/><path d="M14 16h20m-20 8h20m-20 8h20M18 12v24"/></g>`),
		g.Group(children),
	)
}