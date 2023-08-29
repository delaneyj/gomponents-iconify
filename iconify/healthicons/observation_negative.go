package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObservationNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsObservationNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM20.11 23.066h17.956c2.173 0 3.934 1.776 3.934 3.967C42 29.223 40.239 31 38.066 31H38v3h4v2h-3.05a3.5 3.5 0 1 1-4.899 0H13.948a3.5 3.5 0 1 1-4.899 0H6v-2h3V23.343l-.833-.839a4.034 4.034 0 0 1 0-5.676a3.953 3.953 0 0 1 3.892-1.021l.704-.704a3.771 3.771 0 0 1 5.333.002l1.8 1.801a3.77 3.77 0 0 1-.002 5.333l-.368.368l.366.369c.058.058.137.09.218.09Zm-1.993-1.878l.364-.364a1.77 1.77 0 0 0 0-2.504l-1.8-1.8a1.771 1.771 0 0 0-2.504-.002l-.35.35l4.29 4.32ZM36 31v3H11v-8.644l5.275 5.311c.212.213.498.333.797.333H36Zm0-15a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm1-8a1 1 0 1 0-2 0v2.303l.168.252l1 1.5a1 1 0 0 0 1.664-1.11L37 11.697V10Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsObservationNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}