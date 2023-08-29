package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Movie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 7V4H0v9h12v-3l4 2V5l-4 2zm-3 4H2V6h7v5z"/><path fill="currentColor" d="M5 8.4a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}