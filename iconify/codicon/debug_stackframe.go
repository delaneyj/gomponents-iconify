package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStackframe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m14.5 7.15l-4.26-4.74L9.31 2H4.25L3 3.25v9.48l1.25 1.25h5.06l.93-.42l4.26-4.74V7.15zm-5.19 5.58H4.25V3.25h5.06l4.26 4.73l-4.26 4.75z"/>`),
		g.Group(children),
	)
}