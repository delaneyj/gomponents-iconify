package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMonitorCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M20.5 4.5h-15a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h15a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2Zm-16 2a1 1 0 0 1 1-1h15a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-10Z" clip-rule="evenodd"/><path d="M13 17a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/><path fill-rule="evenodd" d="M14.5 17.5h-3a1 1 0 0 0-1 1V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-1-1Zm-3 3.5v-2.5h3V21h-3Z" clip-rule="evenodd"/><path d="M8.5 22a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1h-9Z"/><path fill-rule="evenodd" d="M22 15H4v-1h18v1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMonitorCircleFilled0)"/></g>`),
		g.Group(children),
	)
}