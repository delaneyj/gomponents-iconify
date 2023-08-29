package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4a6 6 0 0 1 6 6v4a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6v-4a6 6 0 0 1 6-6h10Zm-7 5H8v2H6v2h1.999L8 15h2l-.001-2H12v-2h-2V9Zm8 4h-2v2h2v-2Zm-2-4h-2v2h2V9Z"/>`),
		g.Group(children),
	)
}