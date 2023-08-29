package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11a1 1 0 0 0-1 1v7h6v-7a1 1 0 0 0-1-1H3zm1-5V1a1 1 0 1 1 2 0v5h1a1 1 0 0 1 1 1v2.17c1.165.413 2 1.524 2 2.83v7a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-7c0-1.306.835-2.417 2-2.83V7a1 1 0 0 1 1-1h1zm0 2v1h2V8H4z"/>`),
		g.Group(children),
	)
}