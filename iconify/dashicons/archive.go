package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 4v2H1V4h18zM2 7h16v10H2V7zm11 3V9H7v1h6z"/>`),
		g.Group(children),
	)
}