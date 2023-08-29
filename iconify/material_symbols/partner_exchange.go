package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartnerExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9L8.5 5.5L12 2l3.5 3.5L12 9ZM1 20v-4q0-.85.588-1.425T3 14h3.275q.5 0 .95.25t.725.675q.725.975 1.788 1.525T12 17q1.225 0 2.288-.55t1.762-1.525q.325-.425.763-.675t.912-.25H21q.85 0 1.425.575T23 16v4h-7v-2.275q-.875.625-1.888.95T12 19q-1.075 0-2.1-.338T8 17.7V20H1Zm3-7q-1.25 0-2.125-.875T1 10q0-1.275.875-2.138T4 7q1.275 0 2.138.863T7 10q0 1.25-.863 2.125T4 13Zm16 0q-1.25 0-2.125-.875T17 10q0-1.275.875-2.138T20 7q1.275 0 2.138.863T23 10q0 1.25-.863 2.125T20 13Z"/>`),
		g.Group(children),
	)
}