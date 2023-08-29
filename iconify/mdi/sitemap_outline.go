package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SitemapOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16v-3c0-1.11-.89-2-2-2h-6V8h2V2H9v6h2v3H5c-1.11 0-2 .89-2 2v3H1v6h6v-6H5v-3h6v3H9v6h6v-6h-2v-3h6v3h-2v6h6v-6h-2M11 4h2v2h-2V4M5 20H3v-2h2v2m8 0h-2v-2h2v2m8 0h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}