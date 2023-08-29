package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSkipBackOutline0"><g id="evaSkipBackOutline1"><path id="evaSkipBackOutline2" fill="currentColor" d="M16.45 6.2a2.1 2.1 0 0 0-2.21.26l-5.1 4.21l-.14.15V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-3.82l.14.15l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59ZM15.5 16l-4.82-4l4.82-3.91Z"/></g></g>`),
		g.Group(children),
	)
}