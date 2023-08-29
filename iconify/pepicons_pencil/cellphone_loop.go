package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M12.675 15.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M12 14.963a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 3.5A2.5 2.5 0 0 1 8.5 1h7A2.5 2.5 0 0 1 18 3.5v13a2.5 2.5 0 0 1-2.5 2.5H8a2 2 0 0 1-2-2v-1a.5.5 0 0 1 1 0v1a1 1 0 0 0 1 1h7.5a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 15.5 2h-7A1.5 1.5 0 0 0 7 3.5v3a.5.5 0 0 1-1 0v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.64 8.974a2.693 2.693 0 1 0-3.152 4.367A2.693 2.693 0 0 0 7.64 8.974Zm-4.57.022a3.693 3.693 0 1 1 1.258 5.421L1.994 17.65a.5.5 0 1 1-.81-.585l2.333-3.232a3.694 3.694 0 0 1-.447-4.836Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}