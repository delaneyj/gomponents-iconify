package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.437 20.937H5.563a2.372 2.372 0 0 1-2.5-2.211v-11a2.372 2.372 0 0 1 2.5-2.212h.462a.5.5 0 0 1 0 1h-.462a1.381 1.381 0 0 0-1.5 1.212v11a1.38 1.38 0 0 0 1.5 1.211h12.874a1.38 1.38 0 0 0 1.5-1.211v-11a1.381 1.381 0 0 0-1.5-1.212h-.462a.5.5 0 0 1 0-1h.462a2.372 2.372 0 0 1 2.5 2.212v11a2.372 2.372 0 0 1-2.5 2.211Z"/><path fill="currentColor" d="m8.645 6.213l3-3a.5.5 0 0 1 .35-.15a.508.508 0 0 1 .36.15l3 3a.5.5 0 0 1-.71.71l-2.14-2.14v8.47a.508.508 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-8.49l-2.15 2.16a.5.5 0 0 1-.71-.71Z"/>`),
		g.Group(children),
	)
}