package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareChevronsLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-2.293 6.293a1 1 0 0 0-1.414 0l-3 3l-.083.094a1 1 0 0 0 .083 1.32l3 3l.094.083a1 1 0 0 0 1.32-.083l.083-.094a1 1 0 0 0-.083-1.32L14.415 12l2.292-2.293l.083-.094a1 1 0 0 0-.083-1.32zm-5 0a1 1 0 0 0-1.414 0l-3 3l-.083.094a1 1 0 0 0 .083 1.32l3 3l.094.083a1 1 0 0 0 1.32-.083l.083-.094a1 1 0 0 0-.083-1.32L9.415 12l2.292-2.293l.083-.094a1 1 0 0 0-.083-1.32z"/></g>`),
		g.Group(children),
	)
}