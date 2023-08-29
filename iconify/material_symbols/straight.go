package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Straight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21V6.825L9.4 8.4L8 7l4-4l4 4l-1.4 1.4L13 6.825V21h-2Z"/>`),
		g.Group(children),
	)
}