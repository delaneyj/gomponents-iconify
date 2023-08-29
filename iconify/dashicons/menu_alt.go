package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 11h14V9H3v2zm0 5h14v-2H3v2zM3 4v2h14V4H3z"/>`),
		g.Group(children),
	)
}