package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridThreeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v2h2V0H0zm3 0v2h2V0H3zm3 0v2h2V0H6zM0 3v2h2V3H0zm3 0v2h2V3H3zm3 0v2h2V3H6zM0 6v2h2V6H0zm3 0v2h2V6H3zm3 0v2h2V6H6z"/>`),
		g.Group(children),
	)
}