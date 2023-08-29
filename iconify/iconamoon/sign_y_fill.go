package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignYFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.914 10.048C17.734 8.716 16.775 7 15.21 7H14a1 1 0 1 0 0 2h1.21l-2.962 4.814l-1.825-4.867A3 3 0 0 0 7.614 7H7a1 1 0 0 0 0 2h.614a1 1 0 0 1 .936.649l2.37 6.322l-1.998 3.248A1.64 1.64 0 0 1 7.524 20H7a1 1 0 1 0 0 2h.525a3.64 3.64 0 0 0 3.1-1.732l6.289-10.22Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}