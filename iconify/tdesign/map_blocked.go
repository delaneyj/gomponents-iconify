package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapBlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v8h-2v-5l-4 1.667V10.5h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667ZM18 13.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.483 3.483 0 0 0 18 13.5Zm3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745ZM12.5 17a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/>`),
		g.Group(children),
	)
}