package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15h3m-3 0a3 3 0 1 0 3 3v-3m-3 0a3 3 0 1 1 0-6m3 6v-3M9 9h3M9 9a3 3 0 0 1 0-6h3m0 6v3m0-3V3m0 6h3m-3 3a3 3 0 1 0 3-3m-3 3a3 3 0 0 1 3-3m-3-6h3a3 3 0 1 1 0 6"/>`),
		g.Group(children),
	)
}