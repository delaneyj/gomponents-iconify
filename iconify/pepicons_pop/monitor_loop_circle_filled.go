package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorLoopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMonitorLoopCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.5 7.5A2.5 2.5 0 0 1 7 5h11.748a2.5 2.5 0 0 1 2.5 2.5v7.651a2.5 2.5 0 0 1-2.5 2.5h-4.874v2.278a1 1 0 1 1-2 0V17.43a1 1 0 0 1 .626-1.78h6.248a.5.5 0 0 0 .5-.5V7.5a.5.5 0 0 0-.5-.5H7a.5.5 0 0 0-.5.5V9a1 1 0 0 1-2 0V7.5Z"/><path d="M8.187 20.59a1 1 0 0 1 1-1h7.374a1 1 0 1 1 0 2H9.187a1 1 0 0 1-1-1Zm.25-8.208a2.193 2.193 0 1 0-2.058 3.872a2.193 2.193 0 0 0 2.058-3.873Zm-4.732-.032a4.193 4.193 0 1 1 2.674 6.034l-1.677 3.155a1 1 0 1 1-1.766-.939l1.677-3.155a4.194 4.194 0 0 1-.908-5.095Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMonitorLoopCircleFilled0)"/></g>`),
		g.Group(children),
	)
}