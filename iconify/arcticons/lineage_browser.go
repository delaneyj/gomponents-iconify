package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="22.584" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.086 9.24l1.603 4.064a14.213 14.213 0 0 0 8.007 8.007l4.065 1.603a1.168 1.168 0 0 1 0 2.172l-4.065 1.603a14.213 14.213 0 0 0-8.007 8.007l-1.603 4.065a1.168 1.168 0 0 1-2.172 0l-1.603-4.065a14.213 14.213 0 0 0-8.007-8.007l-4.065-1.603a1.168 1.168 0 0 1 0-2.172l4.065-1.603a14.213 14.213 0 0 0 8.007-8.007l1.603-4.065a1.168 1.168 0 0 1 2.172 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.013 18.732h0a10.706 10.706 0 0 0 4.255 4.255h0a1.155 1.155 0 0 1 0 2.026h0a10.706 10.706 0 0 0-4.255 4.255h0a1.155 1.155 0 0 1-2.026 0h0a10.706 10.706 0 0 0-4.255-4.255h0a1.155 1.155 0 0 1 0-2.026h0a10.706 10.706 0 0 0 4.255-4.255h0a1.155 1.155 0 0 1 2.026 0Z"/>`),
		g.Group(children),
	)
}