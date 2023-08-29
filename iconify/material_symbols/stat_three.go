package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.4 21.375l-1.4-1.4l6-6l6 6l-1.4 1.4L12 16.8l-4.6 4.575Zm0-6l-1.4-1.4l6-6l6 6l-1.4 1.4L12 10.8l-4.6 4.575Zm0-6L6 7.975l6-6l6 6l-1.4 1.4L12 4.8L7.4 9.375Z"/>`),
		g.Group(children),
	)
}