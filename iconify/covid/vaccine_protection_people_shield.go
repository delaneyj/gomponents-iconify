package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineProtectionPeopleShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.75 20.893a3.143 3.143 0 1 0 0-6.286a3.143 3.143 0 0 0 0 6.286Zm-.524-8.643h1.048m-.524 0v2.357m3.519-1.116l.74.74m-.37-.37l-1.667 1.667m3.278 1.698v1.048m0-.524h-2.357m1.116 3.519l-.74.74m.37-.37l-1.667-1.667m-1.698 3.278h-1.048m.524 0v-2.357m-3.519 1.116l-.74-.74m.37.37l1.667-1.667m-3.278-1.698v-1.048m0 .524h2.357m-1.116-3.519l.74-.74m-.37.37l1.667 1.667"/><path d="M19.25 8.25V3.76a1.411 1.411 0 0 0-.823-1.292A20.6 20.6 0 0 0 10 .75a20.6 20.6 0 0 0-8.427 1.718A1.411 1.411 0 0 0 .75 3.76v7.224c0 3.532 1.546 8.352 8.228 10.922a2.847 2.847 0 0 0 2.044 0"/><path d="M10.147 9.996a2.791 2.791 0 1 0 0-5.582a2.791 2.791 0 0 0 0 5.582Zm1.72 1.74a4.571 4.571 0 0 0-6.288 4.233"/></g>`),
		g.Group(children),
	)
}