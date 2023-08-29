package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextParagraphTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 2a1 1 0 0 1 1 1v4a3 3 0 0 1-3 3h-4.586l2.293 2.293a1 1 0 0 1-1.414 1.414l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 1 1 1.414 1.414L14.414 8H19a1 1 0 0 0 1-1V3a1 1 0 0 1 1-1ZM7.293 11.293a1 1 0 0 0 0 1.414L9.586 15H3a1 1 0 1 0 0 2h6.586l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4a.997.997 0 0 0 .293-.704V16a.997.997 0 0 0-.295-.71l-3.998-3.997a1 1 0 0 0-1.414 0Z"/>`),
		g.Group(children),
	)
}