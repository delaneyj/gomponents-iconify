package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageStackTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v3h3.293l2.47-2.47a1.75 1.75 0 0 1 2.474 0L13.707 9H17V6a3 3 0 0 0-3-3H6Zm11 7v4a3.001 3.001 0 0 1-2.361 2.932l-3.402-3.402a1.75 1.75 0 0 0-2.474 0L5.36 16.932A3.001 3.001 0 0 1 3 14v-4h14ZM6.707 17h6.586l-2.763-2.763a.75.75 0 0 0-1.06 0L6.707 17Zm1-8h4.586L10.53 7.237a.75.75 0 0 0-1.06 0L7.707 9ZM15.5 5.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}