package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HappyFaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2ZM8.89 13.89a2 2 0 1 1 2 2a2 2 0 0 1-2-2Zm9.24 14.32a8.67 8.67 0 0 1-8.26-6h16.51a8.67 8.67 0 0 1-8.25 6Zm6.93-12.32a2 2 0 1 1 2-2a2 2 0 0 1-2.01 2Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}