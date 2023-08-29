package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareChevronsUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-6.387 10.21a1 1 0 0 0-1.32.083l-3 3l-.083.094a1 1 0 0 0 .083 1.32l.094.083a1 1 0 0 0 1.32-.083L12 14.415l2.293 2.292l.094.083a1 1 0 0 0 1.32-1.497l-3-3zm0-5a1 1 0 0 0-1.32.083l-3 3l-.083.094a1 1 0 0 0 .083 1.32l.094.083a1 1 0 0 0 1.32-.083L12 9.415l2.293 2.292l.094.083a1 1 0 0 0 1.32-1.497l-3-3z"/></g>`),
		g.Group(children),
	)
}