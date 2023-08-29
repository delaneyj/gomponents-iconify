package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 19h-1V9h-4V6.586l-6-6l-6 6V9H2v10H1v2h22v-2ZM6 19H4v-8h2v8Zm12-8h2v8h-2v-8Zm-7 1h2v7h-2v-7Z"/>`),
		g.Group(children),
	)
}