package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxArchiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16l2 4v13a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7.004L4 3Zm9 11v-4h-2v4H8l4 4l4-4h-3Zm6.764-7l-1-2H5.237l-1 2h15.527Z"/>`),
		g.Group(children),
	)
}