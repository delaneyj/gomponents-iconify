package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 6v10h2v2h-2v2h-2v-2h-2v-2h2V6H4V4h12v2zm-4 10h-2v-2h2v2zm6 0h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}