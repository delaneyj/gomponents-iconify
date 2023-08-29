package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 12h2v-2h2V8h-2V6H9v2H7v2h2v2zm1 4c3.9 0 7-3.1 7-7s-3.1-7-7-7s-7 3.1-7 7s3.1 7 7 7zm0-12c2.8 0 5 2.2 5 5s-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5zM3 19h14v-2H3v2z"/>`),
		g.Group(children),
	)
}