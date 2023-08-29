package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm-4 0v-2h2v2H7Zm-4 0V3h2v18H3Zm16-4v-2h2v2h-2Zm-8 0v-2h2v2h-2Zm8-4v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm-4 0v-2h2v2h-2Zm-4 0v-2h2v2H7Zm12-4V7h2v2h-2Zm-8 0V7h2v2h-2Zm8-4V3h2v2h-2Zm-4 0V3h2v2h-2Zm-4 0V3h2v2h-2ZM7 5V3h2v2H7Z"/>`),
		g.Group(children),
	)
}