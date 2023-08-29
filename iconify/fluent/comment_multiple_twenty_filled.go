package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMultipleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2a2.5 2.5 0 0 0-2.456 2.03c.15-.02.302-.03.456-.03h8A3.5 3.5 0 0 1 16 7.5v5.45a2.5 2.5 0 0 0 2-2.45V6a4 4 0 0 0-4-4H6.5Zm-2 3A2.5 2.5 0 0 0 2 7.5v6A2.5 2.5 0 0 0 4.5 16H5v1.028a1 1 0 0 0 1.581.814L9.161 16H12.5a2.5 2.5 0 0 0 2.5-2.5v-6A2.5 2.5 0 0 0 12.5 5h-8Z"/>`),
		g.Group(children),
	)
}