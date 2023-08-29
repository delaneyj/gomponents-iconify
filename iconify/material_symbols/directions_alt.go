package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l8-8q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-8 8q-.3.3-.663.45T12 22Zm0-5l5-5l-5-5l-1.4 1.4l2.55 2.6H7v2h6.15l-2.55 2.6L12 17Z"/>`),
		g.Group(children),
	)
}