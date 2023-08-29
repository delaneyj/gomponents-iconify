package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m31 17l18.795 9.111l15.817 7.668c1.85.897 1.85 3.545 0 4.442L49.795 45.89L31 55"/><path d="m5 17l18.882 9.111l15.891 7.668c1.86.897 1.86 3.545 0 4.442l-15.89 7.668L5 55"/></g>`),
		g.Group(children),
	)
}