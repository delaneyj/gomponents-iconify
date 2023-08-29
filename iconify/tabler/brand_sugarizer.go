package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSugarizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.277 16l3.252-3.252a1.61 1.61 0 0 0-2.277-2.276L12 13.723l-3.252-3.251a1.61 1.61 0 0 0-2.276 2.276L9.723 16l-3.251 3.252a1.61 1.61 0 1 0 2.276 2.277L12 18.277l3.252 3.252a1.61 1.61 0 1 0 2.277-2.277L14.277 16zM9 5a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/>`),
		g.Group(children),
	)
}