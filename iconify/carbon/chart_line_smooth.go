package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineSmooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 15c1.485 0 2.554 1.497 3.686 3.081C17.998 19.918 19.486 22 22 22c5.67 0 7.78-10.79 8-12l-1.968-.358C27.55 12.282 25.394 20 22 20c-1.485 0-2.554-1.497-3.686-3.081C17.002 15.082 15.514 13 13 13c-4.186 0-7.445 7.404-9 11.762V2H2v26a2.002 2.002 0 0 0 2 2h26v-2H5.044c1.51-5.143 4.92-13 7.956-13Z"/>`),
		g.Group(children),
	)
}