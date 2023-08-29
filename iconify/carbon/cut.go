package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.5 19.63L20.24 16l6.26-3.63a5 5 0 0 0-1.21-9.2A5.19 5.19 0 0 0 24 3a5 5 0 0 0-4.33 7.53a5 5 0 0 0 2.39 2.1l-3.82 2.21L4 6.6L3 8.34L16.24 16L3 23.68l1 1.74l14.24-8.26l3.82 2.21a5 5 0 0 0-2.39 2.1A5 5 0 0 0 24 29a5.19 5.19 0 0 0 1.29-.17a5 5 0 0 0 1.21-9.2Zm-5.1-10.1a3 3 0 0 1 1.1-4.12a3 3 0 0 1 4.1 1.11a3 3 0 0 1-1.1 4.11a3 3 0 0 1-4.1-1.1Zm5.2 16a3 3 0 0 1-4.1 1.11a3 3 0 0 1-1.1-4.12a3 3 0 0 1 4.1-1.1a3 3 0 0 1 1.1 4.06Z"/>`),
		g.Group(children),
	)
}