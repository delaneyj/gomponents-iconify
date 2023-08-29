package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.06 9h-26a1 1 0 1 1 0-2h26a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31.06 14h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M31.06 19h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M31.06 24h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M31.06 29h-26a1 1 0 0 1 0-2h26a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M5.56 22.54a1 1 0 0 1-.7-1.71L7.68 18l-2.82-2.83a1 1 0 0 1 0-1.41a1 1 0 0 1 1.41 0L10.51 18l-4.24 4.24a1 1 0 0 1-.71.3Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}