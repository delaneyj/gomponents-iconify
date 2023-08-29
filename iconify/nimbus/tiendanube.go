package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiendanube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.25 2.24a5.79 5.79 0 0 0-4 1.63a4.48 4.48 0 1 0 0 8.26a5.76 5.76 0 1 0 4-9.89zm0 10.24A4.49 4.49 0 0 1 5.76 8H4.48a5.74 5.74 0 0 0 .89 3.07a3.29 3.29 0 0 1-.88.13a3.2 3.2 0 0 1 0-6.4A3.2 3.2 0 0 1 7.69 8H9a4.42 4.42 0 0 0-1.63-3.43a4.48 4.48 0 1 1 2.88 7.91z"/>`),
		g.Group(children),
	)
}