package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListAddFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18h10v2H2v-2Zm0-7h20v2H2v-2Zm0-7h20v2H2V4Zm16 14v-3h2v3h3v2h-3v3h-2v-3h-3v-2h3Z"/>`),
		g.Group(children),
	)
}