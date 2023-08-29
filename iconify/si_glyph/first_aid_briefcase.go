package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidBriefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.82 3H3.18C1.975 3 1 3.99 1 5.209v7.58C1 14.01 1.975 15 3.18 15h11.64c1.205 0 2.18-.99 2.18-2.21V5.208C17.001 3.989 16.026 3 14.82 3zm-4.785 7.016v1.953H8v-1.953H5.987V8H8V5.969h2.035V8h1.952v2.016h-1.952zm1.937-8.047c0-1.038-.839-1.922-1.867-1.922h-2.19c-1.03 0-1.914.884-1.914 1.922l.971-.007c0-.598.363-1.026.956-1.026h2.167c.594 0 .984.429.984 1.026l.893.007z"/>`),
		g.Group(children),
	)
}