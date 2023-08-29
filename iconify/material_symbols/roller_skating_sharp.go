package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerSkatingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16V1h8v3H9v1h3v.5q0 .125.013.25t.037.25H9v1h3.4q.35.575.888.975t1.187.6L20 10.1V16H4Zm1 7q-1.25 0-2.125-.875T2 20q0-1.25.875-2.125T5 17q1.25 0 2.125.875T8 20q0 1.25-.875 2.125T5 23Zm14 0q-1.25 0-2.125-.875T16 20q0-1.25.875-2.125T19 17q1.25 0 2.125.875T22 20q0 1.25-.875 2.125T19 23Zm-7 0q-1.25 0-2.125-.875T9 20q0-1.25.875-2.125T12 17q1.25 0 2.125.875T15 20q0 1.25-.875 2.125T12 23Z"/>`),
		g.Group(children),
	)
}