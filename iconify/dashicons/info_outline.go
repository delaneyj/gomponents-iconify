package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 15h2V9H9v6zm1-10c-.5 0-1 .5-1 1s.5 1 1 1s1-.5 1-1s-.5-1-1-1zm0-4c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9zm0 16c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7z"/>`),
		g.Group(children),
	)
}