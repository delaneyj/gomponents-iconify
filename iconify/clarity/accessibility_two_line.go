package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30.06 11h-24a1 1 0 1 0 0 2H14v9.65l-3.75 10a1 1 0 0 0 .58 1.29a1.13 1.13 0 0 0 .36.06a1 1 0 0 0 .93-.65l3.5-9.35h4.76l3.52 9.35a1 1 0 0 0 .93.65a1.13 1.13 0 0 0 .36-.06a1 1 0 0 0 .58-1.29L22 22.68V13h8.06a1 1 0 1 0 0-2ZM20 22h-4v-9h4Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 10a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}