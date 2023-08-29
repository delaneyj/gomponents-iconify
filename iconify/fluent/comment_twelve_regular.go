package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 3a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h1a.5.5 0 0 1 .5.5v.617L6.743 8.07A.5.5 0 0 1 7 8h1.5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-5ZM2 3.5A1.5 1.5 0 0 1 3.5 2h5A1.5 1.5 0 0 1 10 3.5v4A1.5 1.5 0 0 1 8.5 9H7.138l-2.38 1.429A.5.5 0 0 1 4 10V9h-.5A1.5 1.5 0 0 1 2 7.5v-4Z"/>`),
		g.Group(children),
	)
}