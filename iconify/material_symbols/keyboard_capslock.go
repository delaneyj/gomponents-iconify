package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardCapslock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18v-2h12v2H6Zm6-12.4l6 6l-1.4 1.4L12 8.4L7.4 13L6 11.6l6-6Z"/>`),
		g.Group(children),
	)
}