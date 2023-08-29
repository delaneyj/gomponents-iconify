package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageDotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.019 1.04c-4.398 0-7.968 2.62-7.968 5.852c0 2.95 2.975 5.384 6.842 5.787l-2.048 3.25l6.119-3.603c2.942-.861 5.025-2.968 5.025-5.435c0-3.231-3.569-5.851-7.97-5.851zM6 8H4V6h2v2zm4 0H8V6h2v2zm4 0h-2V6h2v2z"/>`),
		g.Group(children),
	)
}