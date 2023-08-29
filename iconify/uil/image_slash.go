package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 4H10a1 1 0 0 0 0 2h9.5a1 1 0 0 1 1 1v6.76l-1.88-1.88a3 3 0 0 0-1.14-.71a1 1 0 1 0-.64 1.9a.82.82 0 0 1 .36.23l3.31 3.29a.66.66 0 0 0 0 .15a.83.83 0 0 0 0 .15a1.18 1.18 0 0 0 .13.18a.48.48 0 0 0 .09.11a.9.9 0 0 0 .2.14a.6.6 0 0 0 .11.06a.91.91 0 0 0 .37.08a1 1 0 0 0 1-1V7a3 3 0 0 0-2.91-3ZM3.21 2.29a1 1 0 0 0-1.42 1.42L3.18 5.1A3 3 0 0 0 2.5 7v10a3 3 0 0 0 3 3h12.59l1.7 1.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM4.5 7a1 1 0 0 1 .12-.46l2.72 2.71a3 3 0 0 0-1 .63L4.5 11.76Zm1 11a1 1 0 0 1-1-1v-2.42l3.3-3.29a1 1 0 0 1 1.4 0L15.91 18Z"/>`),
		g.Group(children),
	)
}