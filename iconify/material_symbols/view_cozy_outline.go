package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCozyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10.5V3h7.5v7.5H3Zm2-2h3.5V5H5v3.5ZM3 21v-7.5h7.5V21H3Zm2-2h3.5v-3.5H5V19Zm8.5-8.5V3H21v7.5h-7.5Zm2-2H19V5h-3.5v3.5Zm-2 12.5v-7.5H21V21h-7.5Zm2-2H19v-3.5h-3.5V19Zm-7-10.5Zm0 7Zm7-7Zm0 7Z"/>`),
		g.Group(children),
	)
}