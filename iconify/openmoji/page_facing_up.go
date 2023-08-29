package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFacingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="m36 10.958l19.958 20.105v29.895H16.042v-50H36"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m36 10.958l19.958 20.105v29.895H16.042v-50H36"/><path d="m36 10.958l-.042 20.105h13.188"/></g>`),
		g.Group(children),
	)
}