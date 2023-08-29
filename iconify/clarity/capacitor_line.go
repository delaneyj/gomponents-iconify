package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapacitorLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M15 34.06a1 1 0 0 1-1-1V3.15a1 1 0 1 1 2 0v29.91a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M21 34.06a1 1 0 0 1-1-1V3.15a1 1 0 1 1 2 0v29.91a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M14.46 19H3a1 1 0 0 1 0-2h11.46a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M33 19H21.54a1 1 0 0 1 0-2H33a1 1 0 0 1 0 2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}