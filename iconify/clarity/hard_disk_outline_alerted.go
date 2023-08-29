package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDiskOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6 20h24v2H6z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M26 24h4v2h-4z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="m34 21.08l-1.4-5.68h-2.09l1.49 6V29H4v-7.56L7.06 9h12.5l1.15-2H7.06a2 2 0 0 0-1.93 1.47L2 21.08a1 1 0 0 0 0 .24V29a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-7.69a1 1 0 0 0 0-.23Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}