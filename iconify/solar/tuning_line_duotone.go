package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10 9.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm10 5a3 3 0 1 0-6 0a3 3 0 0 0 6 0Z"/><path stroke-linecap="round" d="m17 11l.041-9M7 2v4m0 7v9m10 0v-4" opacity=".5"/></g>`),
		g.Group(children),
	)
}