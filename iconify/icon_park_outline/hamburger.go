package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 22c0-9.941-8.954-18-20-18S4 12.059 4 22h40Z" clip-rule="evenodd"/><path d="M4 38h40v6H4zm0-10l5.455 4l7.272-4L24 32l7.273-4l7.272 4L44 28"/></g>`),
		g.Group(children),
	)
}