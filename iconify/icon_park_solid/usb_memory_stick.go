package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbMemoryStick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUsbMemoryStick0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 22a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V22Zm7-18h18v16H15V4Z"/><path stroke="#000" d="M21 10v2m6-2v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUsbMemoryStick0)"/>`),
		g.Group(children),
	)
}