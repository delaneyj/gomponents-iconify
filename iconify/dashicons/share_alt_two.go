package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m18 8l-5 4V9.01c-2.58.06-4.88.45-7 2.99c.29-3.57 2.66-5.66 7-5.94V3zM4 14h11v-2l2-1.6V16H2V5h9.43c-1.83.32-3.31 1-4.41 2H4v7z"/>`),
		g.Group(children),
	)
}