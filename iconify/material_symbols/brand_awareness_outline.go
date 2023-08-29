package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAwarenessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13v-2h4v2h-4Zm1.2 7L14 17.6l1.2-1.6l3.2 2.4l-1.2 1.6Zm-2-12L14 6.4L17.2 4l1.2 1.6L15.2 8ZM3 15V9h4l5-5v16l-5-5H3Zm7-6.15L7.85 11H5v2h2.85L10 15.15v-6.3ZM7.5 12Z"/>`),
		g.Group(children),
	)
}