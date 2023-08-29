package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 17L24 13C24 8.99999 27 5.99999 31 5.99999C35 5.99999 38 8.99999 38 13L38 32"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 6L10 35C10 39 13 42 17 42C21 42 24 39 24 35L24 17"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 11L10 6L5 11"/><circle cx="38" cy="37" r="5" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}