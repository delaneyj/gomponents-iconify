package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMissedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 21L.4 17.475l1-1.025q2.175-2.225 4.963-3.337T12 12q2.85 0 5.625 1.113T22.6 16.45l1 1.025L19.975 21L16 18v-3.35q-.95-.3-1.95-.475T12 14q-1.05 0-2.05.175T8 14.65V18l-3.975 3Zm7.925-9.65L7 6.4V9H5V3h6v2H8.4l3.525 3.525l5.65-5.65L19 4.3l-7.05 7.05Z"/>`),
		g.Group(children),
	)
}