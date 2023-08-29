package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenhouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L4 9v12h16V9l-8-6m-2 7h4v9h-4v-9m6 0h2v3h-2v-3m-.67-2H8.67L12 5.5L15.33 8M8 10v3H6v-3h2m-2 5h2v4H6v-4m10 4v-4h2v4h-2Z"/>`),
		g.Group(children),
	)
}