package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodFrownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m5.29 10.187l-.424.266l.531.847l.424-.265l-.53-.848Zm4.897-.244l.5.035l.069-.998l-.5-.034l-.069.997Zm-4.366 1.092a7.278 7.278 0 0 1 4.367-1.092l.069-.997a8.278 8.278 0 0 0-4.967 1.241l.53.848ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm-2.5 8A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`),
		g.Group(children),
	)
}