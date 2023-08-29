package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveXCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMoveXCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.026 8.232a1 1 0 0 1 1.409.128l3.333 4a1 1 0 1 1-1.536 1.28l-3.334-4a1 1 0 0 1 .128-1.408Z"/><path d="M16.026 17.768a1 1 0 0 1-.128-1.408l3.334-4a1 1 0 1 1 1.536 1.28l-3.333 4a1 1 0 0 1-1.409.128Z"/><path d="M19 13a1 1 0 0 1-1 1h-8a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1Zm-9.026 4.768a1 1 0 0 1-1.409-.128l-3.333-4a1 1 0 1 1 1.536-1.28l3.334 4a1 1 0 0 1-.128 1.408Z"/><path d="M9.974 8.232a1 1 0 0 1 .128 1.408l-3.334 4a1 1 0 1 1-1.536-1.28l3.333-4a1 1 0 0 1 1.409-.128Z"/><path d="M7 13a1 1 0 0 1 1-1h8a1 1 0 0 1 0 2H8a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMoveXCircleFilled0)"/></g>`),
		g.Group(children),
	)
}