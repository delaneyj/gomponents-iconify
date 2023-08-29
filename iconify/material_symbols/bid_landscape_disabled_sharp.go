package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BidLandscapeDisabledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.8 23.6L18.2 21H3V5.8L.4 3.2l1.425-1.425l20.4 20.4L20.8 23.6Zm.2-5.5l-4.75-4.75l2.75-3.1v-3l-4.175 4.675L5.9 3H21v15.1ZM5 16.95l4-4L13.05 17l.55-.6l-5.45-5.45L5 14.1v2.85Z"/>`),
		g.Group(children),
	)
}