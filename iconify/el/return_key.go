package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReturnKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M808.969 133.929v257.06H942.94v267.899H417.981V508.763L0 787.417l417.982 278.654V915.946H1200V133.928H808.969v.001z"/>`),
		g.Group(children),
	)
}