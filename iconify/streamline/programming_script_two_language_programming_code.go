package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingScriptTwoLanguageProgrammingCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 3.5v8a2 2 0 0 1-2 2H.5a2 2 0 0 0 2-2v-9a2 2 0 0 1 2-2h7.25"/><path d="M11.75.5h0a1.75 1.75 0 0 1 1.75 1.75V3a.5.5 0 0 1-.5.5h-3h0V2.25A1.75 1.75 0 0 1 11.75.5ZM6 4h1.5M5 7h2.5M5 10h2.5"/></g>`),
		g.Group(children),
	)
}