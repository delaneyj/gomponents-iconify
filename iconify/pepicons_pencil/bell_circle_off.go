package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M11 20a.5.5 0 0 1 1 0a1 1 0 1 0 2 0a.5.5 0 0 1 1 0a2 2 0 1 1-4 0Z"/><path fill-rule="evenodd" d="M20.5 17.5a2.96 2.96 0 0 0-1.5-2.575V12a5.5 5.5 0 0 0-5.5-5.5h-1A5.5 5.5 0 0 0 7 12v2.925A2.96 2.96 0 0 0 5.5 17.5a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2ZM18 15.558l.295.133l.055.024A1.96 1.96 0 0 1 19.5 17.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1c0-.769.45-1.467 1.15-1.784l.055-.025l.295-.133V12a4.5 4.5 0 0 1 4.5-4.5h1A4.5 4.5 0 0 1 18 12v3.558Z" clip-rule="evenodd"/><path d="M12.5 4.5a.5.5 0 0 1 1 0V7a.5.5 0 0 1-1 0V4.5Zm-8.35.378a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}