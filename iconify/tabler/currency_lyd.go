package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyLyd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 15h.01M21 5v10a2 2 0 0 1-2 2h-2.764a2 2 0 0 1-1.789-1.106L14 15M5 8l2.773 4.687c.427.697.234 1.626-.43 2.075A1.38 1.38 0 0 1 6.57 15H4.346a.93.93 0 0 1-.673-.293L3 14"/>`),
		g.Group(children),
	)
}