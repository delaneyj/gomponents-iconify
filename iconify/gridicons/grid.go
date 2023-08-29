package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8H4V4h4v4zm6-4h-4v4h4V4zm6 0h-4v4h4V4zM8 10H4v4h4v-4zm6 0h-4v4h4v-4zm6 0h-4v4h4v-4zM8 16H4v4h4v-4zm6 0h-4v4h4v-4zm6 0h-4v4h4v-4z"/>`),
		g.Group(children),
	)
}