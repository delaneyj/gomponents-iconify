package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 2a1 1 0 0 0-1.864-.504l-7 12A1 1 0 0 0 6 15h4v7a1 1 0 0 0 1.864.504l7-12A1 1 0 0 0 18 9h-4V2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}