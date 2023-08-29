package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularjsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l9.3 3.32l-1.418 12.31L12 22l-7.88-4.37L2.7 5.32L12 2Zm0 2.21L6.186 17.26h2.168l1.169-2.92h4.935l1.168 2.92h2.168L12 4.21Zm1.698 8.33h-3.396L12 8.45l1.698 4.09Z"/>`),
		g.Group(children),
	)
}