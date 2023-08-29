package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M42 21H6v21h36V21Z"/><path stroke-linecap="round" d="M14 21V6h20v15m-14-9v2m8-2v2"/></g>`),
		g.Group(children),
	)
}