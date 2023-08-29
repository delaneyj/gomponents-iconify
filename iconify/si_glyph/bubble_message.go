package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.988 1.002c-4.387 0-7.947 2.619-7.947 5.852c0 2.95 2.968 5.383 6.824 5.786l-2.042 3.332l6.103-3.686c2.934-.86 5.012-2.967 5.012-5.433c0-3.232-3.559-5.851-7.95-5.851z"/>`),
		g.Group(children),
	)
}