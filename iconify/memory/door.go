package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 10h2v2h-2v-2m4-8v1h1v15h2v2H3v-2h2V3h1V2h10m-1 2H7v14h8V4Z"/>`),
		g.Group(children),
	)
}