package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubwayLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.2 20l1.8 1.5v.5H5v-.5L6.8 20H5a2 2 0 0 1-2-2V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v11a2 2 0 0 1-2 2h-1.8ZM13 5v6h6V7a2 2 0 0 0-2-2h-4Zm-2 0H7a2 2 0 0 0-2 2v4h6V5Zm8 8H5v5h14v-5ZM7.5 17a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}