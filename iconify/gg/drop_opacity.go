package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOpacity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.945 21.956A9 9 0 0 1 5.635 7.5L12 1.136L18.364 7.5a8.97 8.97 0 0 1 1.991 3.012a9.002 9.002 0 0 1-1.991 9.716a8.987 8.987 0 0 1-2.419 1.728ZM7.05 8.914L12 3.964l4.95 4.95a6.977 6.977 0 0 1 2.048 4.783H5.002A6.976 6.976 0 0 1 7.05 8.914Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}