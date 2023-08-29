package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletListTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h2v2H1zm0 5h2v2H1zm0 5h2v2H1zm5-9h11v1H6zm0 5h11v1H6zm0 5h11v1H6z"/>`),
		g.Group(children),
	)
}