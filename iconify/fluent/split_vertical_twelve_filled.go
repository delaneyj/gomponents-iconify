package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1.5a.5.5 0 0 0-1 0v9a.5.5 0 0 0 1 0v-9Zm-5 2A1.5 1.5 0 0 1 2.5 2H4v8H2.5A1.5 1.5 0 0 1 1 8.5v-5ZM7 10h1.5A1.5 1.5 0 0 0 10 8.5v-5A1.5 1.5 0 0 0 8.5 2H7v8Z"/>`),
		g.Group(children),
	)
}