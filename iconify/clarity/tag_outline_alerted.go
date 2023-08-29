package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="10.52" cy="10.52" r="1.43" fill="currentColor" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="m31.93 19.2l-3.8-3.8h-2.82l5.22 5.22l-9.9 9.9L6 15.91V6h9.92l3.41 3.41l1-1.78l-3-3A2 2 0 0 0 15.92 4H6a2 2 0 0 0-2 2v9.92a2 2 0 0 0 .59 1.41l14.6 14.6a2 2 0 0 0 2.83 0l9.9-9.9a2 2 0 0 0 .01-2.83Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}