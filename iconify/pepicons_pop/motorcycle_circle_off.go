package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorcycleCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M15.75 15.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3.5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M13 11a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="M13 17a2 2 0 0 1 2 2v1.5a2 2 0 1 1-4 0V19a2 2 0 0 1 2-2Z"/><path fill-rule="evenodd" d="M18 14a5 5 0 0 0-10 0v2.5a2.5 2.5 0 0 0 2.5 2.5h5a2.5 2.5 0 0 0 2.5-2.5V14Zm-8 0a3 3 0 0 1 6 0v2.5a.5.5 0 0 1-.5.5h-5a.5.5 0 0 1-.5-.5V14Z" clip-rule="evenodd"/><path d="M18.5 7.5a1 1 0 1 1 0-2h2a1 1 0 1 1 0 2h-2Zm-13 0a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2h-2Z"/><path d="m6.41 7.046l.476-1.455l4.524.863l-.477 1.456l-4.523-.864Zm8.18-.592l.477 1.456l4.523-.864l-.476-1.455l-4.524.863Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}