package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabResearchSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-2.075 0-3.538-1.463T3 17V8H1V2h14v6h-2v3.025q-1.05.675-1.788 1.688T10.175 15H8v-2h3v-1H8v-2h3V8H5v9q0 1.25.875 2.125T8 20q.75 0 1.375-.325t1.05-.9q.2.5.45.95t.6.875q-.675.65-1.563 1.025T8 22Zm0-4v-2h2.025q-.05.5-.013 1t.163 1H8Zm8.5 1q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19Zm5.1 4l-2.7-2.7q-.55.35-1.15.525T16.5 21q-1.875 0-3.187-1.313T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .65-.175 1.25T20.3 18.9l2.7 2.7l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}