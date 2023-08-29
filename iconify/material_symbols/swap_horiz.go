package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapHoriz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 20l-5-5l5-5l1.4 1.425L5.825 14H13v2H5.825L8.4 18.575L7 20Zm10-6l-1.4-1.425L18.175 10H11V8h7.175L15.6 5.425L17 4l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}