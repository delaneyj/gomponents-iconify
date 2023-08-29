package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V10h-2V5.5l-4 1.667V9h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm5.75-4.837h8.5v10.295l-4.247-2.617l-4.253 2.614V11.996Zm2 2v4.715l2.254-1.385l2.246 1.383v-4.713h-4.5Z"/>`),
		g.Group(children),
	)
}