package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorcycleCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M15.75 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3.5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M13 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm.5 12h-1a1.5 1.5 0 0 0-1.5 1.5v2a1.5 1.5 0 0 0 1.5 1.5h1a1.5 1.5 0 0 0 1.5-1.5v-2a1.5 1.5 0 0 0-1.5-1.5ZM12 18.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-2Z" clip-rule="evenodd"/><path d="M18.5 6.75a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5h-2Zm-13 0a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5h-2Z"/><path d="m7.106 6.364l.302-.97l3.698.742l-.302.97l-3.698-.742Zm7.788-.228l.302.97l3.698-.742l-.302-.97l-3.698.742Z"/><path fill-rule="evenodd" d="M17.5 13.5a4.5 4.5 0 1 0-9 0V16a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-2.5Zm-8 0a3.5 3.5 0 1 1 7 0V16a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1v-2.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}