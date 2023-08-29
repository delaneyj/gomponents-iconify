package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 15h-9.59l9.25-9.25a1 1 0 0 0-1.41-1.41L21 13.59V4a1 1 0 0 0-2 0v13h13a1 1 0 0 0 0-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M4 19a1 1 0 0 0 0 2h9.59l-9.26 9.25a1 1 0 1 0 1.41 1.41L15 22.41V32a1 1 0 0 0 2 0V19Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}