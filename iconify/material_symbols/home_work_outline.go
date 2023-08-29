package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeWorkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9h2V7h-2v2Zm0 4h2v-2h-2v2Zm0 4h2v-2h-2v2Zm0 4v-2h4V5h-9v1.4l-2-1.45V3h13v18h-6ZM1 21V11l7-5l7 5v10H9v-5H7v5H1Zm2-2h2v-5h6v5h2v-7L8 8.45L3 12v7Zm14-9Zm-6 9v-5H5v5v-5h6v5Z"/>`),
		g.Group(children),
	)
}