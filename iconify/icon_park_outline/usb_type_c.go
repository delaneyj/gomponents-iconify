package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbTypeC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 15h18c15 0 15 18 0 18H15C0 33 0 15 15 15Zm6 12v-6m6 6v-6m6 6v-6m-18 6v-6m21 3H12"/>`),
		g.Group(children),
	)
}