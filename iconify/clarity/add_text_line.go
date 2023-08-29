package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 21H13a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M12 16a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2H13a1 1 0 0 0-1 1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M27 27H13a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M15.89 9a1 1 0 0 0-1-1H10V3.21a1 1 0 0 0-2 0V8H2.89a1 1 0 0 0 0 2H8v5.21a1 1 0 0 0 2 0V10h4.89a1 1 0 0 0 1-1Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}