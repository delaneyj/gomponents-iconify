package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.152 7.53l5 8a1 1 0 0 0 1.696 0l5-8A1 1 0 0 0 15 6H5a1 1 0 0 0-.848 1.53Zm9.044.47L10 13.113L6.804 8h6.392Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}