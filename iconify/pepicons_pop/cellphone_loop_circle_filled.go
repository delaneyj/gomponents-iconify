package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneLoopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCellphoneLoopCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M15.675 18.137a.675.675 0 1 1-1.35 0a.675.675 0 0 1 1.35 0Z"/><path fill-rule="evenodd" d="M15 17.962a.175.175 0 1 0 0 .35a.175.175 0 0 0 0-.35Zm-1.175.175a1.175 1.175 0 1 1 2.35 0a1.175 1.175 0 0 1-2.35 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.5 6.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3H11A2.5 2.5 0 0 1 8.5 20v-.5a1 1 0 1 1 2 0v.5a.5.5 0 0 0 .5.5h7.5a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-7a1 1 0 0 0-1 1v2.25a1 1 0 0 1-2 0V6.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.347 12.38a2.193 2.193 0 1 0-2.567 3.555a2.193 2.193 0 0 0 2.567-3.555Zm-4.683-.677a4.193 4.193 0 1 1 1.827 6.342L5.4 20.942a1 1 0 1 1-1.622-1.17l2.091-2.898a4.195 4.195 0 0 1-.205-5.17Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCellphoneLoopCircleFilled0)"/></g>`),
		g.Group(children),
	)
}