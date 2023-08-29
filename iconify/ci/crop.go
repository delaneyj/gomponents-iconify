package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 21v-3m0 0H9.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C6 16.48 6 15.92 6 14.8V6m12 12h3M6 6H3m3 0V3m4 3h4.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V14"/>`),
		g.Group(children),
	)
}