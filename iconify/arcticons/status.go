package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Status(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.5 36.5h-25v-25h25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5H28.674l-4-4l-4 4H7.5a2.006 2.006 0 0 0-2 2v33a2.006 2.006 0 0 0 2 2h33a2.006 2.006 0 0 0 2-2v-33a2.006 2.006 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}