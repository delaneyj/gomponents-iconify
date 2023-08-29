package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paristransporttraffic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="15.426" cy="29.948" r="2.471" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.574" cy="29.948" r="2.471" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.843 9.877v22.618a5.388 5.388 0 0 0 5.376 5.377h19.562a5.388 5.388 0 0 0 5.376-5.377V9.877A5.388 5.388 0 0 0 33.781 4.5H14.219a5.388 5.388 0 0 0-5.376 5.377Z"/><rect width="9.371" height="11.714" x="12.872" y="9.531" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".895"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.219 37.924L6.758 43.5m27.023-5.576l7.461 5.576"/><rect width="9.371" height="11.714" x="25.757" y="9.531" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".895"/>`),
		g.Group(children),
	)
}