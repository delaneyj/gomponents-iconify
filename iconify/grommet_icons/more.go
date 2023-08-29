package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func More(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 13v-2h2v2H3Zm8 0v-2h2v2h-2Zm8 0v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}