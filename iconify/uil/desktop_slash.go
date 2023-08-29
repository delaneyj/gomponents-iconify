package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.66 6H18a1 1 0 0 1 1 1v6a.94.94 0 0 1-.14.5a1 1 0 0 0 .31 1.38a.94.94 0 0 0 .53.16a1 1 0 0 0 .84-.46A2.94 2.94 0 0 0 21 13V7a3 3 0 0 0-3-3h-7.34a1 1 0 0 0 0 2Zm11.05 14.29L5.86 4.45L3.71 2.29a1 1 0 0 0-1.42 1.42l1.4 1.39A3 3 0 0 0 3 7v6a3 3 0 0 0 3 3h5v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 .93-.66l2.36 2.37a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM6 14a1 1 0 0 1-1-1V7a1 1 0 0 1 .12-.46L12.59 14Zm7 4v-2h1.59l2 2Z"/>`),
		g.Group(children),
	)
}