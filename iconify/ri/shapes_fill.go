package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l6 10H6l6-10Zm1 12.5h8v8h-8v-8ZM6.75 22a4.75 4.75 0 1 0 0-9.5a4.75 4.75 0 0 0 0 9.5Z"/>`),
		g.Group(children),
	)
}