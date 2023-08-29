package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.939 2.042V.047l-.893.011l-.005 15.88h6.916V2.022l-6.018.02zm1.088 9.979H6.984v-1.053h1.043v1.053zm0-2H6.968V8.984h1.059v1.037zm1.973 2H8.998v-1.068H10v1.068zm.016-2H8.998V8.968h1.018v1.053zm2 2h-1.042v-1.068h1.042v1.068zm-3.989 2H6.968v-1.053h1.059v1.053zm1.989 0H8.998v-1.068h1.018v1.068zm2.015 0h-1.089v-1.068h1.089v1.068zm0-4h-1.073V8.953h1.073v1.068zm.031-2H6.954V3.953h5.108v4.068z"/>`),
		g.Group(children),
	)
}