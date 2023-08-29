package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceAirtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 12a8 8 0 1 0 16 0a8 8 0 0 0-16 0m5 3v.01"/><path d="M15 15a6 6 0 0 0-6-6m3 6a3 3 0 0 0-3-3"/></g>`),
		g.Group(children),
	)
}