package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTwoWayRightBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.807 2.894L14.049.135a.666.666 0 0 0-.944.005l-.01 1.963H4.383c-.738 0-1.336.599-1.336 1.335v8.622l-1.912.012a.666.666 0 0 0 .006.943l2.787 2.786a.663.663 0 0 0 .943.006l2.758-2.758c.26-.26.258-.756-.005-1.018l-1.666.012V5.941c0-.738.272-.983 1.011-.983h6.112l-.01 1.671c.26.259.681.257.942-.005l2.789-2.787a.667.667 0 0 0 .005-.943z"/>`),
		g.Group(children),
	)
}