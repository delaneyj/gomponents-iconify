package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterDramaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-2.3 0-3.9-1.6T1 14.5q0-1.825 1.338-3.163T5.5 10q1.6 0 2.838 1.05T9.9 13.625q.075.375.363.625t.637.25q.5 0 .8-.338t.225-.762q-.35-2.15-1.875-3.575T6.25 8q.45-1.85 2.038-2.925T12 4q2.95 0 4.975 2.025T19 11q1.575 0 2.788 1.4T23 15.5q0 1.875-1.313 3.188T18.5 20h-12Z"/>`),
		g.Group(children),
	)
}