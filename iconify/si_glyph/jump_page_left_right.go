package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpPageLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.958 8.951V7.007H4.979v-1.94c0-.49-3.714 2.265-3.714 2.265a.946.946 0 0 0-.001 1.323s3.715 2.818 3.715 2.293V8.95h1.979zM11.002 7v1.973h2.046V11c0 .455 3.647-2.316 3.647-2.316a.871.871 0 0 0 0-1.261s-3.714-2.78-3.714-2.388L13.048 7h-2.046zM8 0h2v16H8z"/>`),
		g.Group(children),
	)
}