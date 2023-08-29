package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MopedFrontBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36h-37.68a44 44 0 0 0-84.64 0H48a12 12 0 0 0 0 24h37.68a43.89 43.89 0 0 0 9 16.73A68 68 0 0 0 60 136v64a20 20 0 0 0 20 20h16a32 32 0 0 0 64 0h16a20 20 0 0 0 20-20v-64a68 68 0 0 0-34.72-59.27a43.89 43.89 0 0 0 9-16.73H208a12 12 0 0 0 0-24Zm-72 184a8 8 0 0 1-16 0v-56a8 8 0 0 1 16 0Zm36-24h-12v-32a32 32 0 0 0-64 0v32H84v-60a44 44 0 0 1 88 0ZM108 48a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}