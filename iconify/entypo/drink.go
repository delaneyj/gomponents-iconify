package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.538 2.639C17.932 2.094 18 1 18 1H2s.068 1.094.462 1.639L9 11v6H7c-2 0-2 2-2 2h10s0-2-2-2h-2v-6l6.538-8.361zM9.4 6a1.599 1.599 0 1 1 3.2 0a1.6 1.6 0 0 1-3.2 0z"/>`),
		g.Group(children),
	)
}