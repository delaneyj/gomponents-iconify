package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorLoopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M2 4.5a2 2 0 0 1 2-2h11.748a2 2 0 0 1 2 2v7.651a2 2 0 0 1-2 2h-5.374v2.778a.5.5 0 0 1-1 0V14.15H8.5a.5.5 0 0 1 0-1h7.248a1 1 0 0 0 1-1V4.5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1V7a.5.5 0 0 1-1 0V4.5Z"/><path d="M5.687 17.59a.5.5 0 0 1 .5-.5h7.374a.5.5 0 1 1 0 1H6.187a.5.5 0 0 1-.5-.5Zm-.016-8.65a2.693 2.693 0 1 0-2.527 4.756A2.693 2.693 0 0 0 5.67 8.94Zm-4.524.645a3.693 3.693 0 1 1 1.985 5.199l-1.871 3.52a.5.5 0 0 1-.883-.47l1.87-3.52a3.694 3.694 0 0 1-1.101-4.73Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}