package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCheckLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 11.5V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-1l2 2l4-5"/>`),
		g.Group(children),
	)
}