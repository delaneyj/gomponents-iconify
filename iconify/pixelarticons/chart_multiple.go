package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2H1v16h18V2H3zm0 2h14v12H3V4zm18 2v14H5v2h18V6h-2zM7 8H5v6h2V8zm2-2h2v8H9V6zm6 4h-2v4h2v-4z"/>`),
		g.Group(children),
	)
}