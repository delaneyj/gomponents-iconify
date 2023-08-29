package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerVirusShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.319 21.66H2.912A2.162 2.162 0 0 1 .75 19.5V8.685a2.162 2.162 0 0 1 2.162-2.162h7.9a2.162 2.162 0 0 1 2.162 2.162v.747m-8.98-8.315h5.741A1.08 1.08 0 0 1 10.816 2.2v4.323h-7.9V2.2a1.081 1.081 0 0 1 1.078-1.083v0ZM.75 10.848h8.315M.75 17.335h6.487"/><path d="M23.25 15.546a7.5 7.5 0 0 1-5.87 7.337a7.5 7.5 0 0 1-5.869-7.337v-2.2a1.467 1.467 0 0 1 1.467-1.467h8.805a1.467 1.467 0 0 1 1.467 1.467v2.2Zm-5.87-.733v4.402m-2.201-2.201h4.403"/></g>`),
		g.Group(children),
	)
}