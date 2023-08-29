package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm18-9V6H4v5h16Zm-10 7h10v-5H10v5Zm-6 0h4v-5H4v5Z"/>`),
		g.Group(children),
	)
}