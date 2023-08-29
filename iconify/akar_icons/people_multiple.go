package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="7" cy="5" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 12H5.898a3 3 0 0 0-2.976 2.628l-.641 5.124A2 2 0 0 0 4.266 22H9m12.719-2.248l-.64-5.124A3 3 0 0 0 18.101 12h-2.204a3 3 0 0 0-2.976 2.628l-.641 5.124A2 2 0 0 0 14.266 22h5.468a2 2 0 0 0 1.985-2.248Z"/><circle cx="17" cy="5" r="3"/></g>`),
		g.Group(children),
	)
}