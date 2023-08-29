package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11.9876 32L4 24.0062L12 16"/><path d="M36.0124 16L44 23.9939L36 32"/><path d="M4 24L44 24"/></g>`),
		g.Group(children),
	)
}