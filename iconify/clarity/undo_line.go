package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M20.87 11.14h-13l5.56-5.49A1 1 0 0 0 12 4.22l-8 7.91L12 20a1 1 0 0 0 1.41-1.42l-5.55-5.44h13a9.08 9.08 0 0 1 9.13 9a9 9 0 0 1-5 8a1 1 0 0 0 .94 1.86a11 11 0 0 0-5.06-20.82Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}