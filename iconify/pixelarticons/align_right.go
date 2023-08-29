package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5h16v2H4V5zm8 4h8v2h-8V9zm-8 4v2h16v-2H4zm8 4h8v2h-8v-2z"/>`),
		g.Group(children),
	)
}