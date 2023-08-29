package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12a9.97 9.97 0 0 0 .951 4.262l-.93 4.537a1 1 0 0 0 1.18 1.18l4.537-.93c1.294.61 2.74.95 4.262.95c5.523 0 10-4.476 10-10c0-5.522-4.477-10-10-10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}