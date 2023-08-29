package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m17.09 31.58l-15.32-12a2 2 0 0 1 0-3.15l15.32-12a1.93 1.93 0 0 1 2.06-.22A1.77 1.77 0 0 1 20 6v6.7l10.83-8.28a1.93 1.93 0 0 1 2.06-.22A2 2 0 0 1 34 6v24a2 2 0 0 1-1.11 1.79a1.94 1.94 0 0 1-2.06-.22L20 23.31V30a1.77 1.77 0 0 1-.85 1.79a1.94 1.94 0 0 1-2.06-.22ZM32 30l.06-24L18 16.8V6L3 18l15 12V19.2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}