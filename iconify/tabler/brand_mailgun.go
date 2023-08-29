package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandMailgun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="tablerBrandMailgun0" d="M11 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 12a2 2 0 1 0 4 0a9 9 0 1 0-2.987 6.697"/><path d="M7 12a5 5 0 1 0 10 0a5 5 0 1 0-10 0"/><use href="#tablerBrandMailgun0"/><use href="#tablerBrandMailgun0"/></g>`),
		g.Group(children),
	)
}