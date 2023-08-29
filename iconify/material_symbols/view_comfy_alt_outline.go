package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10V2h8v8H2Zm2-2h4V4H4v4ZM2 22v-8h8v8H2Zm2-2h4v-4H4v4Zm10-10V2h8v8h-8Zm2-2h4V4h-4v4Zm-2 14v-8h8v8h-8Zm2-2h4v-4h-4v4ZM8 8Zm0 8Zm8-8Zm0 8Z"/>`),
		g.Group(children),
	)
}