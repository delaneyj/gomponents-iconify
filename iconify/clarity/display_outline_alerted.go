package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M26 32h-1.74a3.61 3.61 0 0 1-1.5-2.52v-1.35h-1.52v1.37a4.2 4.2 0 0 0 .93 2.5h-8.34a4.2 4.2 0 0 0 .93-2.52v-1.35h-1.52v1.37a3.61 3.61 0 0 1-1.5 2.5h-1.8a1 1 0 1 0 0 2h16.12a.92.92 0 0 0 1-1A1 1 0 0 0 26 32Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M33.68 15.4H32V25H4V5h17.87L23 3H3.5A1.5 1.5 0 0 0 2 4.5v21A1.5 1.5 0 0 0 3.5 27h29a1.5 1.5 0 0 0 1.5-1.5V15.38Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M7.7 23V8.76h12l.92-1.6H6.1V23h1.6z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="M26.85 1.14L21.13 11a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}