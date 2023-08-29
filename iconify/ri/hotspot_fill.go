package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotspotFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2v9h7v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h6Zm2 5a2 2 0 0 1 2 2h-2V7Zm0-3a5 5 0 0 1 5 5h-2a3 3 0 0 0-3-3V4Zm0-3a8 8 0 0 1 8 8h-2a6 6 0 0 0-6-6V1Z"/>`),
		g.Group(children),
	)
}