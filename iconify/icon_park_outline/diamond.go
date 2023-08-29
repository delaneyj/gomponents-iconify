package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10.636 5h26.728L45 18.3L24 43L3 18.3L10.636 5Z" clip-rule="evenodd"/><path d="M10.636 5L24 43L37.364 5M3 18.3h42"/><path d="M15.41 18.3L24 5l8.591 13.3"/></g>`),
		g.Group(children),
	)
}