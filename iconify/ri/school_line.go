package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .586l6 6V9h4v10h1v2H1v-2h1V9h4V6.586l6-6ZM18 19h2v-8h-2v8ZM6 11H4v8h2v-8Zm2-3.586V19h3v-7h2v7h3V7.414l-4-4l-4 4Z"/>`),
		g.Group(children),
	)
}