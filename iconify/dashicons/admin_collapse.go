package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminCollapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2.16A7.84 7.84 0 1 1 2.16 10c0-4.33 3.55-7.84 7.84-7.84zm2 11.72V6.12L6.18 9.97z"/>`),
		g.Group(children),
	)
}