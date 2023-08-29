package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineTornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1l11 19L23 3zm-3.47 2l-1.74 3H6.21L4.47 5h15.06zm-9.27 10h3.48L12 18.01L10.26 15zm4.64-2H9.1l-1.74-3h9.27l-1.73 3z"/>`),
		g.Group(children),
	)
}