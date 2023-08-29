package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionSymbolYinYangReligionTaoYinYangTaoismCulture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5S2.93 4.09 7 7c3.5 2.5 0 6.5 0 6.5"/><circle cx="9" cy="4.5" r="1"/><circle cx="5" cy="9.5" r="1"/></g>`),
		g.Group(children),
	)
}