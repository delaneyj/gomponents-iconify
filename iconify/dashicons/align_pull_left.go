package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignPullLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 16V4H3v12h6zm2-7h6V7h-6v2zm0 4h6v-2h-6v2z"/>`),
		g.Group(children),
	)
}