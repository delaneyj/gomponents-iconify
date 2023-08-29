package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18ZM7.416 14H4v5h16v-5h-3.416a5.001 5.001 0 0 1-9.168 0ZM20 5H4v7h5a3 3 0 1 0 6 0h5V5Z"/>`),
		g.Group(children),
	)
}