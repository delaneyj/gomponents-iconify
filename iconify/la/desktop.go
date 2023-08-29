package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 6v18h13v2h-5v2h12v-2h-5v-2h13V6zm2 2h24v14H4z"/>`),
		g.Group(children),
	)
}