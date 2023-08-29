package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNexo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 3l5 3v12l-5 3l-10-6V9l10 6V9l-5-3z"/><path d="M12 6L7 3L2 6v12l5 3l4.7-3.13"/></g>`),
		g.Group(children),
	)
}