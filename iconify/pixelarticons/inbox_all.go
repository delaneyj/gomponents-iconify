package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18v20H3V2zm2 2v4h4v2h6V8h4V4H5zm14 6h-2v2H7v-2H5v4h14v-4zm0 6h-2v2H7v-2H5v4h14v-4z"/>`),
		g.Group(children),
	)
}