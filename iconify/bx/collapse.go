package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collapse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.121 6.465L14 4.344V10h5.656l-2.121-2.121l3.172-3.172l-1.414-1.414zM4.707 3.293L3.293 4.707l3.172 3.172L4.344 10H10V4.344L7.879 6.465zM19.656 14H14v5.656l2.121-2.121l3.172 3.172l1.414-1.414l-3.172-3.172zM6.465 16.121l-3.172 3.172l1.414 1.414l3.172-3.172L10 19.656V14H4.344z"/>`),
		g.Group(children),
	)
}