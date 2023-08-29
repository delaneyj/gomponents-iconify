package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMultipleThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 7.112V7a4 4 0 0 1 4-4h12A6.5 6.5 0 0 1 29 9.5V17a4 4 0 0 1-4 4v.5a4.5 4.5 0 0 1-4.5 4.5h-5.64l-4.403 3.65C9.48 30.462 8 29.767 8 28.497V26H7a4 4 0 0 1-4-4V11.5a4.502 4.502 0 0 1 3.5-4.388ZM8.5 7h12a4.5 4.5 0 0 1 4.5 4.5V19a2 2 0 0 0 2-2V9.5A4.5 4.5 0 0 0 22.5 5h-12a2 2 0 0 0-2 2Zm-1 2A2.5 2.5 0 0 0 5 11.5V22a2 2 0 0 0 2 2h2a1 1 0 0 1 1 1v2.432l3.862-3.202A1 1 0 0 1 14.5 24h6a2.5 2.5 0 0 0 2.5-2.5v-10A2.5 2.5 0 0 0 20.5 9h-13Z"/>`),
		g.Group(children),
	)
}