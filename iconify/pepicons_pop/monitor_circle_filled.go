package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMonitorCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M20.5 4h-15A2.5 2.5 0 0 0 3 6.5v10A2.5 2.5 0 0 0 5.5 19h15a2.5 2.5 0 0 0 2.5-2.5v-10A2.5 2.5 0 0 0 20.5 4ZM5 6.5a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-10Z" clip-rule="evenodd"/><path d="M13 16.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/><path fill-rule="evenodd" d="M14.5 17.5h-3a1 1 0 0 0-1 1V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-1-1Zm-2.5 3V19h2v1.5h-2Z" clip-rule="evenodd"/><path d="M8.5 22.5a1 1 0 1 1 0-2h9a1 1 0 1 1 0 2h-9Z"/><path fill-rule="evenodd" d="M22 15H4v-1h18v1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMonitorCircleFilled0)"/></g>`),
		g.Group(children),
	)
}