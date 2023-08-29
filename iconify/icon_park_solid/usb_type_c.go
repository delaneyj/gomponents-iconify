package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbTypeC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUsbTypeC0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M15 15h18c15 0 15 18 0 18H15C0 33 0 15 15 15Z"/><path stroke="#000" d="M21 27v-6m6 6v-6m6 6v-6m-18 6v-6m21 3H12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUsbTypeC0)"/>`),
		g.Group(children),
	)
}