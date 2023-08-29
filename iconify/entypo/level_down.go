package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 12V5h3v6h10V8l5 4.5l-5 4.5v-3H3a2 2 0 0 1-2-2z"/>`),
		g.Group(children),
	)
}