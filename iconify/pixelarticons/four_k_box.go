package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourKBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4H1v16h22V4H3zm18 2v12H3V6h18zM7 8H5v5h4v3h2V8H9v3H7V8zm8 0h-2v8h2v-3h2v3h2v-3h-2v-2h2V8h-2v3h-2V8z"/>`),
		g.Group(children),
	)
}