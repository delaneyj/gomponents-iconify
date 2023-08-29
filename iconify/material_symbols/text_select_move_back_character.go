package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectMoveBackCharacter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7ZM3 5V3h2v2H3Zm0 16v-2h2v2H3Zm12 0v-2h2V5h-2V3h6v2h-2v14h2v2h-6Zm-8-5l-4-4l4-4l1.4 1.4L6.825 11H14v2H6.825L8.4 14.6L7 16Z"/>`),
		g.Group(children),
	)
}