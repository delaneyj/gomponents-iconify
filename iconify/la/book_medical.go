package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 4C7.355 4 6 5.355 6 7v18c0 1.645 1.355 3 3 3h17V4zm0 2h15v16H9a2.95 2.95 0 0 0-1 .188V7c0-.566.434-1 1-1zm6 4v3h-3v2h3v3h2v-3h3v-2h-3v-3zM9 24h15v2H9c-.566 0-1-.434-1-1c0-.566.434-1 1-1z"/>`),
		g.Group(children),
	)
}