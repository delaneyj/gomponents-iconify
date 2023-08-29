package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2L3 6v1h14V6l-7-4zM5 8l-.2 7h2.5L7 8H5zm4 0l-.2 7h2.5L11 8H9zm4 0l-.2 7h2.5L15 8h-2zM3 18h14v-2H3v2z"/>`),
		g.Group(children),
	)
}