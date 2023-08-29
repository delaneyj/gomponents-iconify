package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h6v2H7.414l3 3L9 10.414l-3-3V10H4V4Zm10 0h6v6h-2V7.414l-3 3L13.586 9l3-3H14V4Zm-3.586 11l-3 3H10v2H4v-6h2v2.586l3-3L10.414 15ZM15 13.586l3 3V14h2v6h-6v-2h2.586l-3-3L15 13.586Z"/>`),
		g.Group(children),
	)
}