package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorOutdent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 4V3H3v1h4zm10 1V3H8v2h9zM7 7H5V5L1 8.5L5 12v-2h2V7zm10 1V6H8v2h9zm-2 3V9H8v2h7zm2 3v-2H8v2h9zM7 14v-1H3v1h4zm4 3v-2H8v2h3z"/>`),
		g.Group(children),
	)
}