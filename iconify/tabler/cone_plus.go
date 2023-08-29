package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18.03 12.022l-5.16-9.515a1 1 0 0 0-1.74 0L3 17.497v.5C3 19.657 7.03 21 12 21c.17 0 .34-.002.508-.005M16 19h6m-3-3v6"/>`),
		g.Group(children),
	)
}