package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerSkatingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16q-.825 0-1.413-.588T4 14V3q0-.825.588-1.413T6 1h4q.825 0 1.413.588T12 3v1H9.5q-.2 0-.35.15T9 4.5q0 .2.15.35T9.5 5H12v.5q0 .125.013.25t.037.25H9.5q-.2 0-.35.15T9 6.5q0 .2.15.35T9.5 7h2.9q.35.575.888.975t1.187.6L17.1 9.3q1.3.35 2.1 1.412t.8 2.413V14q0 .825-.587 1.412T18 16H6Zm-1 7q-1.25 0-2.125-.875T2 20q0-1.25.875-2.125T5 17q1.25 0 2.125.875T8 20q0 1.25-.875 2.125T5 23Zm14 0q-1.25 0-2.125-.875T16 20q0-1.25.875-2.125T19 17q1.25 0 2.125.875T22 20q0 1.25-.875 2.125T19 23Zm-7 0q-1.25 0-2.125-.875T9 20q0-1.25.875-2.125T12 17q1.25 0 2.125.875T15 20q0 1.25-.875 2.125T12 23Z"/>`),
		g.Group(children),
	)
}