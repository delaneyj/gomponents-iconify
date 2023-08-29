package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipCellphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6h6v8H9V6m13-2v12c0 1.11-.89 2-2 2h-4l-4 4l-4-4H4a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2h16a2 2 0 0 1 2 2m-6 1.09C16 4.5 15.5 4 14.86 4H9.14C8.5 4 8 4.5 8 5.09v9.82C8 15.5 8.5 16 9.14 16h5.72c.64 0 1.14-.5 1.14-1.09V5.09Z"/>`),
		g.Group(children),
	)
}