package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20h-2v-3h-3v-2h3v-3h2v3h3v2h-3v3Zm-7-3H2v-2h10v2Zm3-4H2v-2h13v2Zm0-4H2V7h13v2Z"/>`),
		g.Group(children),
	)
}