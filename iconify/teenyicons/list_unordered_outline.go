package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUnorderedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 7.5h11m-15 0h2m2-4h11m-15 0h2m2 8h11m-15 0h2"/>`),
		g.Group(children),
	)
}