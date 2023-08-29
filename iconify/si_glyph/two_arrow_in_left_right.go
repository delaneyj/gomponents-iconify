package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoArrowInLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.125 5.229a.664.664 0 0 0-.946 0a.679.679 0 0 0 0 .954l1.353 1.828H2c-.553 0-1 .444-1 .992s.447.992 1 .992h3.594L4.21 11.812a.673.673 0 0 0 .472 1.15a.667.667 0 0 0 .475-.198l2.819-3.74l-2.851-3.795zm7.696 0a.665.665 0 0 1 .947 0a.677.677 0 0 1 0 .954L12.44 8.011h3.531c.553 0 1 .444 1 .992a.995.995 0 0 1-1 .992h-3.453l1.266 1.817a.675.675 0 0 1-.473 1.15a.665.665 0 0 1-.474-.198l-2.838-3.74l2.822-3.795z"/>`),
		g.Group(children),
	)
}