package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="15" cy="8" r="4" fill="currentColor"/><path fill="currentColor" d="M15 20s8 0 8-2c0-2.4-3.9-5-8-5s-8 2.6-8 5c0 2 8 2 8 2zM6 10V7H4v3H1v2h3v3h2v-3h3v-2z"/>`),
		g.Group(children),
	)
}