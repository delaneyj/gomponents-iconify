package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LedVariantOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a4 4 0 0 0-4 4v6H6v2h3v6h2v-6h2v6h2v-6h3v-2h-2V7a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}