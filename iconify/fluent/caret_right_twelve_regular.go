package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRightTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.994c0-.887 1.07-1.335 1.702-.712l2.037 2.006a1 1 0 0 1 0 1.425L5.702 8.719C5.069 9.34 4 8.893 4 8.006V3.994ZM7.037 6L5 3.994v4.012L7.037 6Z"/>`),
		g.Group(children),
	)
}