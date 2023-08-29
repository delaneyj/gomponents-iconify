package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSocketAu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.22 2A2.22 2.22 0 0 0 2 4.22v15.56C2 21 3 22 4.22 22h15.56A2.22 2.22 0 0 0 22 19.78V4.22C22 3 21 2 19.78 2H4.22M12 4a8 8 0 0 1 8 8a8 8 0 0 1-8 8a8 8 0 0 1-8-8a8 8 0 0 1 8-8M8.27 7.54l-2 3.46L8 12l2-3.46l-1.73-1m7.46 0l-1.73 1L16 12l1.73-1l-2-3.46M11 14v4h2v-4h-2Z"/>`),
		g.Group(children),
	)
}