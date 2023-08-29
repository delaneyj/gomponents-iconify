package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallOrangeDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f6921e" d="M32.725 56.013L8.407 32.077l24.32-23.941l24.315 23.941z"/>`),
		g.Group(children),
	)
}