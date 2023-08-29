package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowchartCircleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 6.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H8v1.793L9.207 12H11v-.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5V13H9.207l-1.353 1.354a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 0 1 0-.708L7 10.793V9h-.5a.5.5 0 0 1-.5-.5v-2ZM2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm8-7a7 7 0 1 0 0 14a7 7 0 0 0 0-14Z"/>`),
		g.Group(children),
	)
}