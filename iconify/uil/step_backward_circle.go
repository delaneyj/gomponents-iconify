package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 7.38a2 2 0 0 0-2 0l-4 2.31V8a1 1 0 0 0-2 0v8a1 1 0 0 0 2 0v-1.69l4 2.31a2 2 0 0 0 2 0a2 2 0 0 0 1-1.73V9.11a2 2 0 0 0-1-1.73Zm-1 7.51L9.5 12l5-2.89ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}