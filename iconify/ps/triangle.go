package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M473 373q7-11 0-21L259 11q-6-10-18-10t-18 10L10 352q-9 10 0 21q3 11 17 11h426q12 0 20-11zM65 341L240 62l175 279H65z"/>`),
		g.Group(children),
	)
}