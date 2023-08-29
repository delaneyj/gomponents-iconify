package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OxygenTankNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsOxygenTankNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM17 6.048h2v-2h-6v2h2v2h-2.17a3.001 3.001 0 1 0 0 2H15v2.083a6.002 6.002 0 0 0-5 5.917v2h12v-2a6.002 6.002 0 0 0-5-5.917v-2.083h5v-2h-5v-2Zm-7 37v-21h12v21a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1Zm-1-34a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm21 8a4 4 0 0 0-4 4v6a4 4 0 0 0 8 0v-6a4 4 0 0 0-4-4Zm-2 4a2 2 0 1 1 4 0v6a2 2 0 0 1-4 0v-6Zm10 5h-3v-2h3a3 3 0 1 1 0 6a1 1 0 0 0-1 1v1h4v2h-5a1 1 0 0 1-1-1v-2a3 3 0 0 1 3-3a1 1 0 0 0 0-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsOxygenTankNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}