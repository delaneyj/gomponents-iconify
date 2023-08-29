package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricBikeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-2.125 0-3.563-1.438T0 12q0-2.125 1.463-3.563T5 7q1.925 0 3.238 1.15T9.9 11h.65l-1.8-5H7V4h5v2h-1.1l.35 1h4.8L14.6 3H12V1h4.025L18.2 6.95h.8q2.075 0 3.538 1.463T24 11.95q0 2.1-1.45 3.575T19 17q-1.8 0-3.163-1.125T14.1 13H9.9q-.35 1.725-1.7 2.863T5 17Zm0-4h2.8v-2H5v2Zm8 10l-6-3h4v-2l6 3h-4v2Zm-.3-12h1.4q.125-.575.338-1.075T15 9h-3.05l.75 2Zm5.3 1.35l1.9-.7l-1-2.65l-1.85.7l.95 2.65Z"/>`),
		g.Group(children),
	)
}