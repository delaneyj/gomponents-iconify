package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutdentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.06 9h-26a1 1 0 1 1 0-2h26a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M31.06 14h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M31.06 19h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M31.06 24h-17a1 1 0 0 1 0-2h17a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M31.06 29h-26a1 1 0 0 1 0-2h26a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M9.56 22.54a1 1 0 0 1-.7-.3L4.61 18l4.25-4.24a1 1 0 0 1 1.41 1.41L7.44 18l2.83 2.83a1 1 0 0 1-.71 1.71Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}