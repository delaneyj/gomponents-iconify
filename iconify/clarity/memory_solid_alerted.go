package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemorySolidAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 17v-1.93h-4V20h-4v-4.93h-1.77a3.68 3.68 0 0 1-2.23-.76V20h-4v-8h2.61A3.68 3.68 0 0 1 19 9.55L20.52 7H4a2 2 0 0 0-2 2v4h2v4H2v10a2 2 0 0 0 2 2h12.61v-3.45H19V29h13a2 2 0 0 0 2-2V17Zm-20 3H8v-8h4Z" class="clr-i-solid--alerted clr-i-solid-path-1--alerted"/><path fill="currentColor" d="m26.85.8l-5.72 9.91a1.28 1.28 0 0 0 1.1 1.91h11.45a1.28 1.28 0 0 0 1.1-1.91L29.06.8a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-solid--alerted clr-i-solid-path-2--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}