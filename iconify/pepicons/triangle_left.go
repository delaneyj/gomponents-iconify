package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.47 4.152l-8 5a1 1 0 0 0 0 1.696l8 5A1 1 0 0 0 14 15V5a1 1 0 0 0-1.53-.848ZM12 6.804v6.392L6.887 10L12 6.804Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}