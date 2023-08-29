package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm13.05 5.588l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508h-7.11v-2h7.11l-2.48-2.508l1.422-1.406ZM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2H8Z"/>`),
		g.Group(children),
	)
}