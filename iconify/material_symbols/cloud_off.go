package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L17.15 20H6.5q-2.3 0-3.9-1.6T1 14.5q0-1.925 1.188-3.425T5.25 9.15q.075-.2.15-.388t.15-.412L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Zm1.8-3.85L8.05 5.225q.875-.6 1.863-.912T12 4q2.925 0 4.963 2.038T19 11q1.725.2 2.863 1.488T23 15.5q0 .975-.375 1.813T21.6 18.75Z"/>`),
		g.Group(children),
	)
}