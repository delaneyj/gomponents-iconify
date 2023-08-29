package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V5zm5 0v7l2.293-2.293a1 1 0 0 1 1.414 0L14 12V5a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}