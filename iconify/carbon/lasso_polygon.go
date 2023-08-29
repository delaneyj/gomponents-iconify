package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LassoPolygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.625 2.22a1 1 0 0 0-1.098-.101L17.935 7.8L3.366 2.07a1 1 0 0 0-1.28 1.336l6.435 14.48a3.965 3.965 0 0 0 1.39 6.944A4.005 4.005 0 0 1 6 28H4v2h2a6.004 6.004 0 0 0 5.928-5.12a4.002 4.002 0 0 0 2.93-2.88h8.382a2 2 0 0 0 1.927-1.465l4.796-17.267a1 1 0 0 0-.338-1.049ZM11 23a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm12.24-3h-8.382a3.897 3.897 0 0 0-4.512-2.934L4.905 4.824L18.065 10l9.345-5.012Z"/>`),
		g.Group(children),
	)
}