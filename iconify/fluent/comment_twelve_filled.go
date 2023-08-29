package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v4A1.5 1.5 0 0 0 3.5 9H4v1a.5.5 0 0 0 .757.429L7.138 9H8.5A1.5 1.5 0 0 0 10 7.5v-4A1.5 1.5 0 0 0 8.5 2h-5Z"/>`),
		g.Group(children),
	)
}