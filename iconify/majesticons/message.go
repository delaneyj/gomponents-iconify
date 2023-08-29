package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Message(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3a1 1 0 0 1 1 1v.97c0 1.659 1.904 2.596 3.22 1.584l4.35-3.347a1 1 0 0 1 .61-.207H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}