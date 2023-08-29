package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightGalleryTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6Zm-1 5h-4V8h4v3Zm-4 1h4v1a2 2 0 0 1-2 2h-2v-3Zm4-6v1h-4V4h2a2 2 0 0 1 2 2Zm-5-2v11H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h7Z"/>`),
		g.Group(children),
	)
}