package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 3h1v10H0V3zm8 0h2v10H8V3zm3 0h1v10h-1V3zm2 0h1v10h-1V3zm2 0h1v10h-1V3zM2 3h3v10H2V3zm4 0h1v10H6V3z"/>`),
		g.Group(children),
	)
}