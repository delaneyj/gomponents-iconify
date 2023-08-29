package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.4 22L6 20.6l6-6l6 6l-1.4 1.4l-4.6-4.6L7.4 22ZM12 9.4l-6-6L7.4 2L12 6.6L16.6 2L18 3.4l-6 6Z"/>`),
		g.Group(children),
	)
}