package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M20.5 4h-15A2.5 2.5 0 0 0 3 6.5v10A2.5 2.5 0 0 0 5.5 19h15a2.5 2.5 0 0 0 2.5-2.5v-10A2.5 2.5 0 0 0 20.5 4ZM5 6.5a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-10Z" clip-rule="evenodd"/><path d="M13 16.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/><path fill-rule="evenodd" d="M14.5 17.5h-3a1 1 0 0 0-1 1V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-1-1Zm-2.5 3V19h2v1.5h-2Z" clip-rule="evenodd"/><path d="M8.5 22.5a1 1 0 1 1 0-2h9a1 1 0 1 1 0 2h-9Z"/><path fill-rule="evenodd" d="M22 15H4v-1h18v1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}