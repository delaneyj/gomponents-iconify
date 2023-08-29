package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalleryThumbnailSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h14v14H1Zm16-8V5h6v6h-6ZM4 15h8l-2.625-3.5L7.5 14l-1.375-1.825L4 15Zm13 4v-6h6v6h-6Z"/>`),
		g.Group(children),
	)
}