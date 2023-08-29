package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M13.71 12.59a1 1 0 0 0-1.39-.26l-6.53 4.45a1 1 0 0 0 0 1.65l6.53 4.45a1 1 0 1 0 1.13-1.65l-5.32-3.62L13.45 14a1 1 0 0 0 .26-1.41Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m30.21 16.78l-6.53-4.45A1 1 0 1 0 22.55 14l5.32 3.63l-5.32 3.63a1 1 0 0 0 1.13 1.65l6.53-4.45a1 1 0 0 0 0-1.65Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M19.94 9.83a.9.9 0 0 0-1.09.66l-3.44 13.8a.9.9 0 0 0 .66 1.09h.22a.9.9 0 0 0 .87-.68l3.44-13.81a.9.9 0 0 0-.66-1.06Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}