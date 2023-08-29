package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDrupal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2c0 4.308-7 6-7 12a7 7 0 0 0 14 0c0-6-7-7.697-7-12z"/><path d="M12 11.33a65.753 65.753 0 0 1-2.012 2.023C8.988 14.31 8 15.32 8 17c0 2.17 1.79 4 4 4s4-1.827 4-4c0-1.676-.989-2.685-1.983-3.642c-.42-.404-2.259-2.357-5.517-5.858l3.5 3.83z"/></g>`),
		g.Group(children),
	)
}