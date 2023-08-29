package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerStandby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M312 87.666v33.47C381.676 144.582 432 210.522 432 288c0 97.047-78.953 176-176 176S80 385.047 80 288c0-77.478 50.324-143.418 120-166.864v-33.47C112.422 112.179 48 192.7 48 288c0 114.691 93.309 208 208 208s208-93.309 208-208c0-95.3-64.422-175.821-152-200.334Z"/><path fill="currentColor" d="M240 16h32v272h-32z"/>`),
		g.Group(children),
	)
}