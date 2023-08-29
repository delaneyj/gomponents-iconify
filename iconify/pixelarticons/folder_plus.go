package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h8v2h10v14H2V4h2zm16 4H10V6H4v12h16V8zm-6 2h2v2h2v2h-2v2h-2v-2h-2v-2h2v-2z"/>`),
		g.Group(children),
	)
}