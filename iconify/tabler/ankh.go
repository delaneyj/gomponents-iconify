package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ankh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 13h12m-6 8v-8l-.422-.211A6.472 6.472 0 0 1 8 7a4 4 0 1 1 8 0a6.472 6.472 0 0 1-3.578 5.789L12 13"/>`),
		g.Group(children),
	)
}