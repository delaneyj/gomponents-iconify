package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildArrowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M24.82 15.8a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.41L27.2 21H9V3.78a1 1 0 1 0-2 0V21a2 2 0 0 0 2 2h18.15l-3.74 3.75a1 1 0 0 0 0 1.41a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29L31 22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}