package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorcycleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMotorcycleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M15.75 15.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3.5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M13 11a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="M13 17a2 2 0 0 1 2 2v1.5a2 2 0 1 1-4 0V19a2 2 0 0 1 2-2Z"/><path fill-rule="evenodd" d="M18 14a5 5 0 0 0-10 0v2.5a2.5 2.5 0 0 0 2.5 2.5h5a2.5 2.5 0 0 0 2.5-2.5V14Zm-8 0a3 3 0 0 1 6 0v2.5a.5.5 0 0 1-.5.5h-5a.5.5 0 0 1-.5-.5V14Z" clip-rule="evenodd"/><path d="M18.5 7.5a1 1 0 1 1 0-2h2a1 1 0 1 1 0 2h-2Zm-13 0a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2h-2Z"/><path d="m6.41 7.046l.476-1.455l4.524.863l-.477 1.456l-4.523-.864Zm8.18-.592l.477 1.456l4.523-.864l-.476-1.455l-4.524.863Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMotorcycleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}