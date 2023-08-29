package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.4 18.4L6 17l6-6l6 6l-1.4 1.4l-4.6-4.575L7.4 18.4Zm0-6L6 11l6-6l6 6l-1.4 1.4L12 7.825L7.4 12.4Z"/>`),
		g.Group(children),
	)
}