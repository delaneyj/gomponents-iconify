package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbMicroOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUsbMicroOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 18a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V18Z"/><path fill="#555" d="M11 15h26v8H11v-8Z"/><path d="M21 23v-2m6 2v-2m5 2v-2m-16 2v-2m-8-6h32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUsbMicroOne0)"/>`),
		g.Group(children),
	)
}