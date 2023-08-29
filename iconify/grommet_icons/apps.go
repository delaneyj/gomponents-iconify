package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 5h2V3h-2v2Zm-8 0h2V3h-2v2ZM3 5h2V3H3v2Zm16 8h2v-2h-2v2Zm-8 0h2v-2h-2v2Zm-8 0h2v-2H3v2Zm16 8h2v-2h-2v2Zm-8 0h2v-2h-2v2Zm-8 0h2v-2H3v2Z"/>`),
		g.Group(children),
	)
}