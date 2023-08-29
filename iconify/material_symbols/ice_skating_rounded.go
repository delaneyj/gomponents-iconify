package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceSkatingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h3v-2H5q-.825 0-1.413-.588T3 16V5q0-.825.588-1.413T5 3h4q.825 0 1.413.588T11 5v1H8.5q-.2 0-.35.15T8 6.5q0 .2.15.35T8.5 7h2.55v1H8.5q-.2 0-.35.15T8 8.5q0 .2.15.35T8.5 9h2.9q.35.575.888.975t1.187.6l2.625.725q1.3.35 2.1 1.412t.8 2.413V16q0 .825-.587 1.412T17 18h-1v2h2q1.25 0 2.125-.875T21 17q0-.425.288-.713T22 16q.425 0 .713.288T23 17q0 2.075-1.463 3.538T18 22H3Zm5-2h6v-2H8v2Z"/>`),
		g.Group(children),
	)
}