package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUsbOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M42 21H6v21h36V21Z"/><path stroke-linecap="round" d="M14 21V6h20v15m-14-9v2m8-2v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUsbOne0)"/>`),
		g.Group(children),
	)
}