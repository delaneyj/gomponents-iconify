package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ccw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.25 15.5a8.766 8.766 0 0 1-8.75 8.75a8.7 8.7 0 0 1-6.366-2.764l2.068-1.442l-7.9-3.703l.743 8.695l2.193-1.53a12.217 12.217 0 0 0 9.26 4.243c6.767 0 12.25-5.482 12.25-12.25h-3.5zM15.5 6.75a8.704 8.704 0 0 1 6.366 2.764l-2.068 1.443l7.9 3.7l-.745-8.692l-2.192 1.53A12.22 12.22 0 0 0 15.5 3.248C8.733 3.25 3.25 8.733 3.25 15.5h3.5a8.76 8.76 0 0 1 8.75-8.75z"/>`),
		g.Group(children),
	)
}