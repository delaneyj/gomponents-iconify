package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerVirusBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.408 21.443H2.962a2.177 2.177 0 0 1-2.178-2.178V8.374A2.178 2.178 0 0 1 2.962 6.2h7.757A2.178 2.178 0 0 1 12.9 8.374v.956"/><path d="M4.052.75H9.63a1.089 1.089 0 0 1 1.089 1.089V6.2H2.962V1.839A1.089 1.089 0 0 1 4.052.75v0Zm10.66 18.891a3.461 3.461 0 0 1 4.894-4.895m1.014 2.447a3.46 3.46 0 0 1-3.461 3.461m-.577-9.517h1.154m-.577 0v2.596m6.057 2.884v1.153m0-.577H20.62m1.23 3.875l-.816.816m.408-.408l-1.836-1.835m-1.87 3.609h-1.154m.577 0v-2.596m-6.056-2.884v-1.153m0 .576h2.595m-1.229-3.874l.815-.816m-.407.408l1.835 1.835m8.504-3.609L11.103 23.25M.784 10.552h8.58m-8.58 6.535h6.535"/></g>`),
		g.Group(children),
	)
}