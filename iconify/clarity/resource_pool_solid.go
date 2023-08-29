package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResourcePoolSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8.57 30.9A16 16 0 0 0 33.95 19H18.43Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M33.95 17A16 16 0 1 0 7 29.6L17.49 17Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}