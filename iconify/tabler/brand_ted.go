package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8h4M4 8v8m9-8H9v8h4m-4-4h2.5M16 8v8h2a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3h-2z"/>`),
		g.Group(children),
	)
}