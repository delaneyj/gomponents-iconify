package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.1 13h5.8l1.74-3H7.36zm2.9 5.01L13.74 15h-3.48zM4.47 5l1.74 3h11.58l1.74-3z" opacity=".3"/><path fill="currentColor" d="m1 3l11 19L23 3H1zm11 15.01L10.26 15h3.48L12 18.01zM14.9 13H9.1l-1.74-3h9.27l-1.73 3zM6.21 8L4.47 5h15.06l-1.74 3H6.21z"/>`),
		g.Group(children),
	)
}