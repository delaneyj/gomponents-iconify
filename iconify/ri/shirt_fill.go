package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.998 4v7l5-2.5l5 2.5V4h3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3Zm5 4l-4.5-5h9l-4.5 5Zm1 3.236l-1-.5l-1 .5V20h2v-8.764Zm2 2.764v2h4v-2h-4Z"/>`),
		g.Group(children),
	)
}