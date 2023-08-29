package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V6a2 2 0 0 1 4 0v6m0-2a2 2 0 1 1 2 2h-6m0 0h6a2 2 0 0 1 0 4h-6m2 0a2 2 0 1 1-2 2v-6m0 0v6a2 2 0 0 1-4 0v-6m0 2a2 2 0 1 1-2-2h6m0 0H6a2 2 0 0 1 0-4h6m-2 0a2 2 0 1 1 2-2v6"/>`),
		g.Group(children),
	)
}