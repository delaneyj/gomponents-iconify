package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceNintendo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 20V4H7a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h3zm4 0V4h3a4 4 0 0 1 4 4v8a4 4 0 0 1-4 4h-3z"/><circle cx="17.5" cy="15.5" r="1" fill="currentColor"/><circle cx="6.5" cy="8.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}