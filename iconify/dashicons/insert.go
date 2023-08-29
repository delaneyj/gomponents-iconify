package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 1c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9zm0 16c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7zm1-11H9v3H6v2h3v3h2v-3h3V9h-3V6z" class="st0"/>`),
		g.Group(children),
	)
}