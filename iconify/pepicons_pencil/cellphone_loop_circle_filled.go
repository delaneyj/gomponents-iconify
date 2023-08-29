package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCellphoneLoopCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M15.675 18.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M15 17.962a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 6.5A2.5 2.5 0 0 1 11.5 4h7A2.5 2.5 0 0 1 21 6.5v13a2.5 2.5 0 0 1-2.5 2.5H11a2 2 0 0 1-2-2v-1a.5.5 0 0 1 1 0v1a1 1 0 0 0 1 1h7.5a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 18.5 5h-7A1.5 1.5 0 0 0 10 6.5v3a.5.5 0 0 1-1 0v-3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.64 11.974a2.693 2.693 0 1 0-3.152 4.367a2.693 2.693 0 0 0 3.152-4.367Zm-4.57.022a3.693 3.693 0 1 1 1.258 5.421L4.994 20.65a.5.5 0 1 1-.81-.585l2.333-3.232a3.694 3.694 0 0 1-.447-4.836Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCellphoneLoopCircleFilled0)"/></g>`),
		g.Group(children),
	)
}