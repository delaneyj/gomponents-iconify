package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10l-2 7h4l-2.925 4.225L14.65 11.8l.55-.8h-1.35l-1.175-1.175L14.35 4H9v2.15l-2-2V2Zm3 20v-8H7V9.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L13.75 16.6L10 22Zm1.825-13.025Z"/>`),
		g.Group(children),
	)
}