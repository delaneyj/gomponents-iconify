package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentBanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M7 43h32L23 17L7 43Z"/><path d="M18.077 25L23 28l4.923-3M23 17V4m12 0H23v8h12l-3-4l3-4Z"/></g>`),
		g.Group(children),
	)
}