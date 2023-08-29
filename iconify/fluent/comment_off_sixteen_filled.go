package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.293 12l2.853 2.854a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l.739.738A2.495 2.495 0 0 0 1 4.5v5A2.5 2.5 0 0 0 3.5 12H4v1.942a.98.98 0 0 0 1.625.738L8.688 12h2.605ZM15 9.5c0 .916-.492 1.716-1.227 2.152L4.121 2H12.5A2.5 2.5 0 0 1 15 4.5v5Z"/>`),
		g.Group(children),
	)
}