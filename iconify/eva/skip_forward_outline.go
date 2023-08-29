package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaSkipForwardOutline0"><g id="evaSkipForwardOutline1"><path id="evaSkipForwardOutline2" fill="currentColor" d="M16 6a1 1 0 0 0-1 1v3.82l-.14-.15l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21l.14-.15V17a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1Zm-7.5 9.91V8l4.82 4Z"/></g></g>`),
		g.Group(children),
	)
}