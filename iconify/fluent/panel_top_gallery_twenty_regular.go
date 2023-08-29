package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelTopGalleryTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H5ZM3 7a2 2 0 0 1 2-2h2v4H3V7Zm0 3h14v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3Zm14-1h-4V5h2a2 2 0 0 1 2 2v2Zm-5-4v4H8V5h4Z"/>`),
		g.Group(children),
	)
}