package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8 34a1 1 0 0 1-1-1V2.92a1 1 0 0 1 2 0V33a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M17 33.92a1 1 0 0 1-1-1V9.1a1 1 0 1 1 2 0v23.82a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M26 34a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v24a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M33.11 18h-25a1 1 0 1 1 0-2h25a1 1 0 1 1 0 2Z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M33.1 26.94h-25a1 1 0 1 1 0-1.94h25a1 1 0 1 1 0 1.92Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="currentColor" d="M33 8.92H3A1 1 0 1 1 3 7h30a1 1 0 1 1 0 1.94Z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}