package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 14.5a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm-10-5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/><path stroke-linecap="round" d="M16.959 9V2M6.958 15v7m10.001 0v-2M6.958 2v2"/></g>`),
		g.Group(children),
	)
}