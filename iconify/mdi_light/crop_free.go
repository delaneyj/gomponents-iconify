package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4h2v1H6a2 2 0 0 0-2 2v2H3V7a3 3 0 0 1 3-3M4 18a2 2 0 0 0 2 2h2v1H6a3 3 0 0 1-3-3v-2h1v2M17 4a3 3 0 0 1 3 3v2h-1V7a2 2 0 0 0-2-2h-2V4h2m3 14a3 3 0 0 1-3 3h-2v-1h2a2 2 0 0 0 2-2v-2h1v2Z"/>`),
		g.Group(children),
	)
}