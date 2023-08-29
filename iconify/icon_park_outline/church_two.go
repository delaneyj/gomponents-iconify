package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="m13 24l11-10l11 10v20H13V24ZM7 8l3-4l3 4v36H7V8Zm28 0l3-4l3 4v36h-6V8Z"/><path d="M24 25v10m-4-6h8"/></g>`),
		g.Group(children),
	)
}