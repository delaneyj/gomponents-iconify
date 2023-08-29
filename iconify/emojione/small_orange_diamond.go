package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallOrangeDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f27a52" d="m16.019 32l15.98-15.981l15.98 15.98L32 47.98z"/>`),
		g.Group(children),
	)
}