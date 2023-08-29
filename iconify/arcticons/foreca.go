package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foreca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.008" cy="23.899" r="17.83" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.124 6.23a3.305 3.305 0 0 1 1.427-3.19A3.189 3.189 0 0 1 26 2.974a3.302 3.302 0 0 1 1.544 3.132M11.265 10.85a3.168 3.168 0 0 1-.456-3.307a3.155 3.155 0 0 1 5.655-.226m15.324.411a3.024 3.024 0 1 1 4.901 3.311M6.432 20.106a3.384 3.384 0 0 1-2.414-2.68a3.361 3.361 0 0 1 1.391-3.317a3.425 3.425 0 0 1 3.618-.192m30.428.203a3.12 3.12 0 0 1 3.305.207a3.091 3.091 0 0 1-.988 5.497M6.318 26.317L3.363 30.05l4.463 1.753m2.667 3.863l.061 5.103l4.817-1.054m5.65 1.809l3.202 3.976l3.217-4.02m5.315-2.062l4.654 1.39l.29-5.415m2.487-3.805l4.45-1.498l-2.638-4.364"/>`),
		g.Group(children),
	)
}