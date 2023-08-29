package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 34.15a15 15 0 0 1-7.52-28a1 1 0 0 1 1 1.73a13 13 0 1 0 13 0a1 1 0 1 1 1-1.73a15 15 0 0 1-7.52 28Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}