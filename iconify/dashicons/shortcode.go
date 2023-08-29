package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shortcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 14H4V6h2V4H2v12h4zm1.1 3h2.1l3.7-14h-2.1zM14 4v2h2v8h-2v2h4V4z"/>`),
		g.Group(children),
	)
}