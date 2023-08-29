package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g transform="translate(3 3)"><path fill-rule="evenodd" d="M10 13.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z" clip-rule="evenodd"/><rect width="1" height="3" x="1" y="10.5" rx=".5" transform="rotate(-90 1 10.5)"/><rect width="1" height="3" x="16" y="10.5" rx=".5" transform="rotate(-90 16 10.5)"/><rect width="1" height="3" x="14" y="14.707" rx=".5" transform="rotate(-45 14 14.707)"/><rect width="1" height="3" x="3" y="3.707" rx=".5" transform="rotate(-45 3 3.707)"/><rect width="1" height="3" x="9.5" y="16" rx=".5"/><rect width="1" height="3" x="9.5" y="1" rx=".5"/><rect width="1" height="3" x="16.121" y="3" rx=".5" transform="rotate(45 16.121 3)"/><rect width="1" height="3" x="5.121" y="14" rx=".5" transform="rotate(45 5.121 14)"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}