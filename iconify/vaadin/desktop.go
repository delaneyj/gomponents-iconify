package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 0H0v13h6v2H4v1h8v-1h-2v-2h6V0zM9 12H7v-1h2v1zm6-2H1V1h14v9z"/>`),
		g.Group(children),
	)
}