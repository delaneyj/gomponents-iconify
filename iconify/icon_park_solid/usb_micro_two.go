package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbMicroTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUsbMicroTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 18a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v6.14a3 3 0 0 1-.456 1.59l-3.663 5.86A3 3 0 0 1 37.337 33H10.663a3 3 0 0 1-2.544-1.41l-3.663-5.86A3 3 0 0 1 4 24.14V18Z"/><path fill="#fff" stroke="#000" d="M11 15h26v8H11v-8Z"/><path stroke="#000" d="M21 23v-2m6 2v-2m5 2v-2m-16 2v-2"/><path stroke="#fff" d="M8 15h32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUsbMicroTwo0)"/>`),
		g.Group(children),
	)
}