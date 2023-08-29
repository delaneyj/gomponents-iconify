package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorLoopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMonitorLoopCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 7.5a2 2 0 0 1 2-2h11.748a2 2 0 0 1 2 2v7.651a2 2 0 0 1-2 2h-5.374v2.778a.5.5 0 0 1-1 0V17.15H11.5a.5.5 0 0 1 0-1h7.248a1 1 0 0 0 1-1V7.5a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1V10a.5.5 0 0 1-1 0V7.5Z"/><path d="M8.687 20.59a.5.5 0 0 1 .5-.5h7.374a.5.5 0 1 1 0 1H9.187a.5.5 0 0 1-.5-.5Zm-.016-8.65a2.693 2.693 0 1 0-2.527 4.756A2.693 2.693 0 0 0 8.67 11.94Zm-4.524.645a3.693 3.693 0 1 1 1.985 5.199l-1.871 3.52a.5.5 0 0 1-.883-.47l1.87-3.52a3.694 3.694 0 0 1-1.101-4.73Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMonitorLoopCircleFilled0)"/></g>`),
		g.Group(children),
	)
}