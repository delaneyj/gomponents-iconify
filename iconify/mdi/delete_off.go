package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5.27L3.28 4L5 5.72l.28.28l1 1L18 18.72l2 2L18.73 22l-1.46-1.46c-.34.29-.77.46-1.27.46H8c-1.1 0-2-.9-2-2V9.27l-4-4M19 4v2H7.82l-2-2H8.5l1-1h5l1 1H19m-1 3v9.18L8.82 7H18Z"/>`),
		g.Group(children),
	)
}