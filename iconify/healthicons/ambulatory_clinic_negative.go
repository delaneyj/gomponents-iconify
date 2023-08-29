package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmbulatoryClinicNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsAmbulatoryClinicNegative0)"><path d="M26 40h8V24h-8v16Zm-5-16h-7v7h7v-7ZM20 9h-2v2h-2v2h2v2h2v-2h2v-2h-2V9Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm24.653 9.985H42.5V12.5L38 15.532V40h3a1 1 0 1 1 0 2h-3v.015H10V42H7a1 1 0 1 1 0-2h3V15.497l-4-3.03V9.984h7.347a6.002 6.002 0 0 1 11.306 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAmbulatoryClinicNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}