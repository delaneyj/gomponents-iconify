package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="m4 8l8-4l8 4l-8 4l-8-4z"/><path stroke-linecap="round" d="m4 12l8 4l8-4M4 16l8 4l8-4"/></g>`),
		g.Group(children),
	)
}