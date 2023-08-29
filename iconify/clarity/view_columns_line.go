package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 5H5a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM13 29H5V7h8Zm10 0h-8V7h8Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}