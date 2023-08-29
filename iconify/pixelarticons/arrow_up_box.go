package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h18V3H3v18zM19 5v14H5V5h14zm-8 12h2v-6h2V9h-2V7h-2v2H9v2h2v6zm-2-4v-2H7v2h2zm8 0h-2v-2h2v2z"/>`),
		g.Group(children),
	)
}