package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Option(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 11a1 1 0 0 0 2 0v-1H5a1 1 0 0 0-1 1z"/><path fill="currentColor" d="M0 0v16h16V0H0zm11 9a2 2 0 1 1-2 2v-1H7v1a2 2 0 1 1-2-2h1V7H5a2 2 0 1 1 2-2v1h2V5a2 2 0 1 1 2 2h-1v2h1z"/><path fill="currentColor" d="M12 5a1 1 0 0 0-2 0v1h1a1 1 0 0 0 1-1zM5 4a1 1 0 0 0 0 2h1V5a1 1 0 0 0-1-1zm2 3h2v2H7V7zm3 4a1 1 0 1 0 1-1h-1v1z"/>`),
		g.Group(children),
	)
}