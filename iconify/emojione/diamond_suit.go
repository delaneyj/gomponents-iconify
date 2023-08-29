package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ff5a79" d="M2.018 32L32 2.02L61.98 32L32 61.981z"/>`),
		g.Group(children),
	)
}