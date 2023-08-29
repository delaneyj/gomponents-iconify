package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalANegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsClinicalANegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-1v2.273a1.5 1.5 0 0 1 1 1.415V39a5 5 0 0 1-10 0V12.687a1.5 1.5 0 0 1 1-1.414V9h-1a1 1 0 0 1-1-1V5Zm4 4v2.688a1.5 1.5 0 0 1-1 1.414V18h6v-4.898a1.5 1.5 0 0 1-1-1.415V9h-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsClinicalANegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}