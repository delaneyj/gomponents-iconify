package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m235.31 89.36l-68.68-68.68a16 16 0 0 0-22.63 0l-28.44 28.44l-58 21.76a16 16 0 0 0-10.2 12.35l-20.77 124.6a4 4 0 0 0 6.77 3.49l57-57a23.85 23.85 0 0 1-2.29-12.08a24 24 0 1 1 13.6 23.4l-57 57a4 4 0 0 0 3.49 6.77l124.61-20.77a16 16 0 0 0 12.35-10.16l21.77-58.07L235.31 112a16 16 0 0 0 0-22.63ZM200 124.68L131.32 56l24-24L224 100.68Z"/>`),
		g.Group(children),
	)
}