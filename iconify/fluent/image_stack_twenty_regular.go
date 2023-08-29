package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageStackTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Zm7.707 13l-2.47-2.47a1.75 1.75 0 0 0-2.474 0L6.293 16H6a2 2 0 0 1-2-2v-4h12v4a2 2 0 0 1-2 2h-.293Zm-6 0l1.763-1.763a.75.75 0 0 1 1.06 0L12.293 16H7.707ZM16 9h-2.793l-1.97-1.97a1.75 1.75 0 0 0-2.474 0L6.793 9H4V6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v3ZM8.207 9L9.47 7.737a.75.75 0 0 1 1.06 0L11.793 9H8.207ZM15 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 7.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}