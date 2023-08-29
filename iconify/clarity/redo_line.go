package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M24 4.22a1 1 0 0 0-1.41 1.42l5.56 5.49h-13A11 11 0 0 0 10.07 32a1 1 0 0 0 .93-1.82a9 9 0 0 1-5-8a9.08 9.08 0 0 1 9.13-9h13l-5.54 5.48A1 1 0 0 0 24 20l8-7.91Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}