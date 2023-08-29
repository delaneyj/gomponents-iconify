package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 6a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V6Zm12 0h-11v7h11V6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.25 5.5a.5.5 0 0 1 .5-.5h14.5a.5.5 0 0 1 0 1H5.75a.5.5 0 0 1-.5-.5Zm2.24 7.598L6.11 20h13.78l-1.38-6.902l.98-.196l1.38 6.902A1 1 0 0 1 19.89 21H6.11a1 1 0 0 1-.98-1.196l1.38-6.902l.98.196Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}