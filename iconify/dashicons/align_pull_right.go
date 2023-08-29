package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignPullRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 16V4h-6v12h6zM9 7H3v2h6V7zm0 4H3v2h6v-2z"/>`),
		g.Group(children),
	)
}