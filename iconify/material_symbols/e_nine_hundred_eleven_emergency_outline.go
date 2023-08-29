package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ENineHundredElevenEmergencyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-2h1.6l1.975-6.575q.2-.65.738-1.038T10.5 10h3q.65 0 1.188.388t.737 1.037L17.4 18H19v2H5Zm3.7-2h6.6l-1.8-6h-3l-1.8 6ZM11 8V3h2v5h-2Zm5.95 2.475L15.525 9.05l3.55-3.525l1.4 1.4l-3.525 3.55ZM18 15v-2h5v2h-5ZM7.05 10.475l-3.525-3.55l1.4-1.4l3.55 3.525l-1.425 1.425ZM1 15v-2h5v2H1Zm11 3Z"/>`),
		g.Group(children),
	)
}