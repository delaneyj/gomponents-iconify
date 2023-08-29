package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelAltMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5H6v10h12v-2h2v-2h2V9h-2V7h-2V5H8zm10 2v2h2v2h-2v2H8V7h10zM4 9H2v10h12v-2H4V9z"/>`),
		g.Group(children),
	)
}