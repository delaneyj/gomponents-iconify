package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fingernail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4"/><path d="m38 9.472l.343 1.056h1.11l-.898.652l.343 1.056l-.898-.652l-.898.652l.343-1.056l-.898-.652h1.11L38 9.472Z"/><rect width="12" height="24" x="18" y="13" stroke-linecap="round" stroke-linejoin="round" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 25c-2 0-5 2.118-5 6v9.784M30 25c2 0 5 2.118 5 6v9.5"/></g>`),
		g.Group(children),
	)
}