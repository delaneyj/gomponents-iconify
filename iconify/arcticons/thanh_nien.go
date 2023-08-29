package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThanhNien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 12.694V5.5h37v7.194h-37Zm32.889 0L24 22.972L9.611 12.694"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 14.236l.835 1.692l1.867.271l-1.351 1.317l.319 1.86l-1.67-.879l-1.67.878l.32-1.86l-1.351-1.316l1.867-.271l.834-1.692ZM8.789 34.278h5.447M11.512 42.5v-8.222m5.602 8.222v-8.222L22.56 42.5v-8.222M33.764 42.5v-8.222L39.21 42.5v-8.222m-8.324 0L28.162 42.5l-2.723-8.222"/>`),
		g.Group(children),
	)
}