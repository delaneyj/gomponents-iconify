package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20V7l-2-4H4L2 7.004V20a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1ZM4 9h16v10H4V9Zm1.236-4h13.528l1 2H4.237l.999-2ZM15 11H9v2h6v-2Z"/>`),
		g.Group(children),
	)
}