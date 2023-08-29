package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v8h-2v-5l-4 1.667V9.5h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm10-6.102l2.21 3.226l3.752 1.106l-2.385 3.1l.108 3.909L18 20.762l-3.685 1.31l.108-3.91l-2.385-3.1l3.751-1.105L18 10.731Zm0 3.538l-.963 1.405l-1.634.482l1.039 1.35l-.047 1.703L18 18.64l1.605.57l-.047-1.703l1.04-1.35l-1.635-.482L18 14.27Z"/>`),
		g.Group(children),
	)
}