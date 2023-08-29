package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Migrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 6h6V4H2v12.01h8V14H4V6zm2 2h6V5l6 5l-6 5v-3H6V8z"/>`),
		g.Group(children),
	)
}