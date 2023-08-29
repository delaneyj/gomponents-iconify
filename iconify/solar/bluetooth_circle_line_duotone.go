package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothCircleLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m11 12l3.2-2.407c.533-.401.8-.602.8-.875c0-.274-.267-.475-.8-.876l-1.454-1.094c-.762-.573-1.143-.86-1.444-.708C11 6.191 11 6.669 11 7.623V12Zm0 0v4.377c0 .954 0 1.432.302 1.583c.301.151.682-.135 1.444-.708l1.454-1.094c.533-.402.8-.602.8-.876c0-.273-.267-.474-.8-.875L11 12Zm0 0L8 9.5m3 2.5l-3 2.5" opacity=".5"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}