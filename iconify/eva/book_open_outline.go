package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBookOpenOutline0"><g id="evaBookOpenOutline1"><path id="evaBookOpenOutline2" fill="currentColor" d="M20.62 4.22a1 1 0 0 0-.84-.2L12 5.77L4.22 4A1 1 0 0 0 3 5v12.2a1 1 0 0 0 .78 1l8 1.8h.44l8-1.8a1 1 0 0 0 .78-1V5a1 1 0 0 0-.38-.78ZM5 6.25l6 1.35v10.15L5 16.4ZM19 16.4l-6 1.35V7.6l6-1.35Z"/></g></g>`),
		g.Group(children),
	)
}