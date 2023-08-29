package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.174 12.594h3.652L12 8.095l-1.826 4.499Z"/><path fill="currentColor" d="M12 1L2 4.652l1.525 13.541L12 23l8.475-4.807L22 4.652L12 1Zm6.24 16.786h-2.33l-1.257-3.212H9.347L8.09 17.786H5.76L12 3.431l6.24 14.355Z"/>`),
		g.Group(children),
	)
}