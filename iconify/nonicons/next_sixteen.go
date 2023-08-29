package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.5 6.5 0 1 0 3.65 11.88L6 6.577v4.673a.75.75 0 0 1-1.5 0V4.5a.75.75 0 0 1 1.327-.48l7.474 9a.75.75 0 0 1-.088 1.049a8 8 0 1 1 2.12-2.865a.75.75 0 1 1-1.375-.602A6.453 6.453 0 0 0 14.5 8A6.5 6.5 0 0 0 8 1.5Zm2.5 2.25a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-1.5 0v-3a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}