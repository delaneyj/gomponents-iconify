package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3h-4v2h4v4h2V5a2 2 0 0 0-2-2m0 16h-4v2h4a2 2 0 0 0 2-2v-4h-2M5 15H3v4a2 2 0 0 0 2 2h4v-2H5M3 5v4h2V5h4V3H5a2 2 0 0 0-2 2Z"/>`),
		g.Group(children),
	)
}