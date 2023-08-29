package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetEnergy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M5 9v6.74C5 19.199 8.134 22 12 22s7-2.802 7-6.26V9M5 9h3m8 0h3"/><path stroke-linecap="round" d="M14.074 11.5v3.56c0 1.072-.928 1.94-2.074 1.94s-2.074-.868-2.074-1.94V11.5"/><path d="M10 13H5m14 0h-5"/><path stroke-linecap="round" d="M11.667 2L10 5h4l-1.667 3"/></g>`),
		g.Group(children),
	)
}