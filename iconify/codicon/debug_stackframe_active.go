package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStackframeActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0z"/><path d="m14.5 7.15l-4.26-4.74L9.31 2H4.25L3 3.25v9.48l1.25 1.25h5.06l.93-.42l4.26-4.74V7.15zm-5.19 5.58H4.25V3.25h5.06l4.26 4.73l-4.26 4.75z"/></g>`),
		g.Group(children),
	)
}