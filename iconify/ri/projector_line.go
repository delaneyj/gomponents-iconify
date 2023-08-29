package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm8.126 9H4v7h16v-7h-1.126a4.002 4.002 0 0 1-7.748 0Zm0-2a4.002 4.002 0 0 1 7.748 0H20V5H4v5h7.126ZM15 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-9 2h2v2H6v-2Z"/>`),
		g.Group(children),
	)
}