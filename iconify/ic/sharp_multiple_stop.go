package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMultipleStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 4l4 4l-4 4V9h-4V7h4V4zM7 17h4v-2H7v-3l-4 4l4 4v-3zm12-2h-2v2h2v-2zm-4 0h-2v2h2v-2zm-4-8H9v2h2V7zM7 7H5v2h2V7z"/>`),
		g.Group(children),
	)
}