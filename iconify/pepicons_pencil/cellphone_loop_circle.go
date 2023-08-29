package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoopCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M15.675 18.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M15 17.962a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 6.5A2.5 2.5 0 0 1 11.5 4h7A2.5 2.5 0 0 1 21 6.5v13a2.5 2.5 0 0 1-2.5 2.5H11a2 2 0 0 1-2-2v-1a.5.5 0 0 1 1 0v1a1 1 0 0 0 1 1h7.5a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 18.5 5h-7A1.5 1.5 0 0 0 10 6.5v3a.5.5 0 0 1-1 0v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.64 11.974a2.693 2.693 0 1 0-3.152 4.367a2.693 2.693 0 0 0 3.152-4.367Zm-4.57.022a3.693 3.693 0 1 1 1.258 5.421L4.994 20.65a.5.5 0 1 1-.81-.585l2.333-3.232a3.694 3.694 0 0 1-.447-4.836Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}