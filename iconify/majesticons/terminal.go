package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm4.617 1.076a1 1 0 0 1 1.09.217l4 4a1 1 0 0 1 0 1.414l-4 4A1 1 0 0 1 6 16V8a1 1 0 0 1 .617-.924zM12 16a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}