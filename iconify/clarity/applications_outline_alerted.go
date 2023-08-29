package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationsOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8 8H4v2h6V4H8v4z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M8 19H4v2h6v-6H8v4z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M19 19h-4v2h6v-6h-2v4z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="M30 15v4h-4v2h6v-6h-2z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted"/><path fill="currentColor" d="M8 30H4v2h6v-6H8v4z" class="clr-i-outline--alerted clr-i-outline-path-5--alerted"/><path fill="currentColor" d="M19 30h-4v2h6v-6h-2v4z" class="clr-i-outline--alerted clr-i-outline-path-6--alerted"/><path fill="currentColor" d="M30 30h-4v2h6v-6h-2v4z" class="clr-i-outline--alerted clr-i-outline-path-7--alerted"/><path fill="currentColor" d="M19 8h-4v2h4v-.11l2-3.39V4h-2Z" class="clr-i-outline--alerted clr-i-outline-path-8--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-9--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}