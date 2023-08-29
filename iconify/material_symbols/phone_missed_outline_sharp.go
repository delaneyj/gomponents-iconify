package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMissedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.025 21L.4 17.475l1-1.025q2.125-2.275 4.938-3.362T12 12q2.85 0 5.638 1.1t4.962 3.35l1 1.025L19.975 21L16 18v-3.35q-.95-.3-1.95-.475T12 14q-1.05 0-2.05.175T8 14.65V18l-3.975 3ZM6 15.45q-.725.375-1.4.863T3.2 17.4l1 1L6 17v-1.55Zm12 .05V17l1.8 1.4l1-.95q-.725-.65-1.4-1.125T18 15.5Zm-12-.05Zm12 .05Zm-6.05-4.15L7 6.4V9H5V3h6v2H8.4l3.525 3.525l5.65-5.65L19 4.3l-7.05 7.05Z"/>`),
		g.Group(children),
	)
}