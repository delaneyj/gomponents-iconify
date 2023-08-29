package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4.5V7H5v2h11v2.5L19.5 8M16 12.5V15H5v2h11v2.5l3.5-3.5"/>`),
		g.Group(children),
	)
}