package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalleryThumbnailOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h14v14H1Zm16-8V5h6v6h-6Zm2-2h2V7h-2v2ZM3 17h10V7H3v10Zm1-2h8l-2.625-3.5L7.5 14l-1.375-1.825L4 15Zm13 4v-6h6v6h-6Zm2-2h2v-2h-2v2ZM3 17V7v10Zm16-8V7v2Zm0 8v-2v2Z"/>`),
		g.Group(children),
	)
}