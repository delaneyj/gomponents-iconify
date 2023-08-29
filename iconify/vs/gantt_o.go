package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GanttO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1408 106v250q0 18-9 30t-23 12H32q-14 0-23-12t-9-30V106q0-18 9-30t23-12h1344q14 0 23 12t9 30zm512 400v250q0 18-9 30t-23 12H544q-14 0-23-12t-9-30V506q0-16 9-28t23-14h1344q14 2 23 14t9 28zm-768 402v250q0 17-9.5 29.5T1120 1200H288q-12 0-22-13t-10-29V908q0-18 10-30t22-12h832q14 0 23 12t9 30zm512 400v250q0 17-9.5 29.5T1632 1600H800q-12 0-22-13t-10-29v-250q0-16 10-29t22-13h832q13 0 22.5 12.5t9.5 29.5z"/>`),
		g.Group(children),
	)
}