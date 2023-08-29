package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxArchiveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 3l2 4v13a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7.004L4 3h16Zm0 6H4v10h16V9Zm-7 1v4h3l-4 4l-4-4h3v-4h2Zm5.764-5H5.237l-1 2h15.527l-1-2Z"/>`),
		g.Group(children),
	)
}