package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 7c-3.302 0-6 2.698-6 6a6 6 0 0 0 2.47 4.844C2.83 19.154 1 21.864 1 25h2c0-3.326 2.674-6 6-6s6 2.674 6 6h2c0-3.326 2.674-6 6-6s6 2.674 6 6h2c0-3.136-1.83-5.846-4.47-7.156A5.998 5.998 0 0 0 29 13c0-3.302-2.698-6-6-6s-6 2.698-6 6a6 6 0 0 0 2.47 4.844a8.048 8.048 0 0 0-3.47 3.28a8.048 8.048 0 0 0-3.47-3.28A5.998 5.998 0 0 0 15 13c0-3.302-2.698-6-6-6zm0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4zm14 0c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4z"/>`),
		g.Group(children),
	)
}