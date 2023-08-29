package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorLoopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M17.248 3.75H5.5A1.5 1.5 0 0 0 4 5.25v7.651a1.5 1.5 0 0 0 1.5 1.5h11.748a1.5 1.5 0 0 0 1.5-1.5V5.25a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill-rule="evenodd" d="M3 5.25a2.5 2.5 0 0 1 2.5-2.5h11.748a2.5 2.5 0 0 1 2.5 2.5v7.651a2.5 2.5 0 0 1-2.5 2.5h-4.874v2.278a1 1 0 1 1-2 0V15.4H5.5A2.5 2.5 0 0 1 3 12.9V5.25Zm2.5-.5a.5.5 0 0 0-.5.5v7.651a.5.5 0 0 0 .5.5h11.748a.5.5 0 0 0 .5-.5V5.25a.5.5 0 0 0-.5-.5H5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.687 18.34a1 1 0 0 1 1-1h7.374a1 1 0 1 1 0 2H7.687a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M8.727 13.566a3.193 3.193 0 1 1-5.639-2.997a3.193 3.193 0 0 1 5.639 2.997Z"/><path fill-rule="evenodd" d="M6.937 10.132a2.193 2.193 0 1 0-2.058 3.872a2.193 2.193 0 0 0 2.058-3.872ZM2.205 10.1a4.193 4.193 0 1 1 2.674 6.034l-1.677 3.155a1 1 0 1 1-1.766-.939l1.677-3.155a4.194 4.194 0 0 1-.908-5.095Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M2 4.5a2 2 0 0 1 2-2h11.748a2 2 0 0 1 2 2v7.651a2 2 0 0 1-2 2h-5.374v2.778a.5.5 0 0 1-1 0V14.15H8.5a.5.5 0 0 1 0-1h7.248a1 1 0 0 0 1-1V4.5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1V7a.5.5 0 0 1-1 0V4.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.687 17.59a.5.5 0 0 1 .5-.5h7.374a.5.5 0 1 1 0 1H6.187a.5.5 0 0 1-.5-.5Zm-.016-8.65a2.693 2.693 0 1 0-2.527 4.756A2.693 2.693 0 0 0 5.67 8.94Zm-4.524.645a3.693 3.693 0 1 1 1.985 5.199l-1.871 3.52a.5.5 0 0 1-.883-.47l1.87-3.52a3.694 3.694 0 0 1-1.101-4.73Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}