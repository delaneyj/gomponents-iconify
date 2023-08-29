package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadConeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.462 1a1 1 0 0 0-.97.757L5.056 7.5H8.5a.5.5 0 0 1 0 1H4.806l-.5 2H9.25a.5.5 0 0 1 0 1H4.056L3.43 14H1.5a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-1.931L9.509 1.757A1 1 0 0 0 8.537 1H7.462Z"/>`),
		g.Group(children),
	)
}