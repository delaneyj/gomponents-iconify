package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 9v7h-3v-6H6v3L1 8.5L6 4v3h11c1.104 0 2 .897 2 2z"/>`),
		g.Group(children),
	)
}