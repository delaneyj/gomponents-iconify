package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlightlySmilingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8H9Zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5H15Zm-6.966 5.134a1 1 0 0 1 1.367.365A2.992 2.992 0 0 0 12 15c1.11 0 2.08-.601 2.6-1.5a1 1 0 1 1 1.73 1A4.998 4.998 0 0 1 12 17a4.998 4.998 0 0 1-4.33-2.5a1 1 0 0 1 .364-1.366Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}