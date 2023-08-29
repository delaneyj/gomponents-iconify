package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M19 4a1 1 0 0 0 0 2h9.59l-9.25 9.25a1 1 0 1 0 1.41 1.41L30 7.41V17a1 1 0 0 0 2 0V4Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M4 19a1 1 0 0 1 2 0v9.59l9.25-9.25a1 1 0 1 1 1.41 1.41L7.41 30H17a1 1 0 0 1 0 2H4Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}