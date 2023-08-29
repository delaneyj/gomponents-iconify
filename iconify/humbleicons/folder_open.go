package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M17 8V7a1 1 0 0 0-1-1h-5.586a1 1 0 0 1-.707-.293L8.293 4.293A1 1 0 0 0 7.586 4H4a1 1 0 0 0-1 1v12a2 2 0 0 0 2 2h14a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v6a2 2 0 0 1-2 2"/>`),
		g.Group(children),
	)
}