package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandElastic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2a5 5 0 0 1 5 5c0 .712-.232 1.387-.5 2c1.894.042 3.5 1.595 3.5 3.5c0 1.869-1.656 3.4-3.5 3.5c.333.625.5 1.125.5 1.5a2.5 2.5 0 0 1-2.5 2.5c-.787 0-1.542-.432-2-1c-.786 1.73-2.476 3-4.5 3a5 5 0 0 1-4.583-7a3.5 3.5 0 0 1-.11-6.992h.195a2.5 2.5 0 0 1 2-4c.787 0 1.542.432 2 1c.786-1.73 2.476-3 4.5-3zM8.5 9l-3-1"/><path d="m9.5 5l-1 4l1 2l5 2l4-4m-.001 7l-3-.5l-1-2.5m.001 6l1-3.5M5.417 15L9.5 11"/></g>`),
		g.Group(children),
	)
}