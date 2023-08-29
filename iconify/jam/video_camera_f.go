package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.98 3.605L16 1.585A2 2 0 0 1 17.414 1H18a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-.586A2 2 0 0 1 16 10.414l-2.02-2.019A4 4 0 0 1 10 12H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h6a4 4 0 0 1 3.98 3.605zM5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}