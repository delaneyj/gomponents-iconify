package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 18a.5.5 0 0 1-.5-.5V16H8a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h2V9H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h4V2.5a.5.5 0 0 1 1 0V4h4a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-4v2h2a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-2v1.5a.5.5 0 0 1-.5.5ZM6 5a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6Zm2 7a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1H8Z"/>`),
		g.Group(children),
	)
}