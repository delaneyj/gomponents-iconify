package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.437 20.948H5.563a2.372 2.372 0 0 1-2.5-2.21v-11a2.372 2.372 0 0 1 2.5-2.211h.462a.5.5 0 0 1 0 1h-.462a1.38 1.38 0 0 0-1.5 1.211v11a1.38 1.38 0 0 0 1.5 1.21h12.874a1.38 1.38 0 0 0 1.5-1.21v-11a1.38 1.38 0 0 0-1.5-1.211h-.462a.5.5 0 0 1 0-1h.462a2.372 2.372 0 0 1 2.5 2.211v11a2.372 2.372 0 0 1-2.5 2.21Z"/><path fill="currentColor" d="m15.355 10.592l-3 3a.5.5 0 0 1-.35.15a.508.508 0 0 1-.36-.15l-3-3a.5.5 0 0 1 .71-.71l2.14 2.139V3.552a.508.508 0 0 1 .5-.5a.5.5 0 0 1 .5.5v8.49l2.15-2.16a.5.5 0 0 1 .71.71Z"/>`),
		g.Group(children),
	)
}