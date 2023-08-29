package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorArrowSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 12.36L3.88 3a1 1 0 0 0-1.29 1.28L12 29.44a1 1 0 0 0 1.89-.05l2.69-8.75l9.12 8.9a1 1 0 0 0 1.41 0l2.35-2.35a1 1 0 0 0 0-1.41l-9.09-8.86L29 14.25a1 1 0 0 0 0-1.89Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}