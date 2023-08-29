package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="m21 12l17 30H4l17-30Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M36.5 42H44l-8-14l-3 5M21 12l-8 30"/></g>`),
		g.Group(children),
	)
}