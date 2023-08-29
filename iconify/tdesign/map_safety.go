package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V10h-2V5.5l-4 1.667V9h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm5.498-4.835h9v5.633a3 3 0 0 1-1.36 2.512l-3.14 2.052l-3.14-2.052a3 3 0 0 1-1.36-2.512v-5.633Zm2 2v3.633a1 1 0 0 0 .453.837l2.047 1.337l2.047-1.337a1 1 0 0 0 .453-.837v-3.633h-5Z"/>`),
		g.Group(children),
	)
}