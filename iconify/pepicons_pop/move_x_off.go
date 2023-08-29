package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveXOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.027 5.232a1 1 0 0 1 1.408.128l3.333 4a1 1 0 1 1-1.536 1.28l-3.334-4a1 1 0 0 1 .129-1.408Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.027 14.768a1 1 0 0 1-.129-1.408l3.334-4a1 1 0 1 1 1.536 1.28l-3.333 4a1 1 0 0 1-1.409.128Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16 10a1 1 0 0 1-1 1H7a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1Zm-9.026 4.768a1 1 0 0 1-1.409-.128l-3.333-4a1 1 0 1 1 1.536-1.28l3.334 4a1 1 0 0 1-.128 1.408Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.974 5.232a1 1 0 0 1 .128 1.408l-3.334 4a1 1 0 1 1-1.536-1.28l3.333-4a1 1 0 0 1 1.409-.128Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4 10a1 1 0 0 1 1-1h8a1 1 0 0 1 0 2H5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}