package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4.008h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2ZM19 8a1 1 0 0 0-1 1v5.392l-1.48-1.48a2.78 2.78 0 0 0-3.929 0l-.698.697l-2.486-2.486a2.777 2.777 0 0 0-3.924 0L4 12.606V7a1.001 1.001 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3.003 3.003 0 0 0-3 3v12a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3V9a1 1 0 0 0-1-1ZM5 20a1.001 1.001 0 0 1-1-1v-3.566l2.897-2.897a.8.8 0 0 1 1.096 0l3.168 3.167c.009.01.012.022.02.03L15.448 20Zm13-1a.971.971 0 0 1-.179.537l-4.514-4.514l.698-.698a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		g.Group(children),
	)
}