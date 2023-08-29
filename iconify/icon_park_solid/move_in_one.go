package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveInOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMoveInOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" d="m28 28l14 2.8l-4.2 2.8l4.2 4.2l-4.2 4.2l-4.2-4.2l-2.8 4.2L28 28Z" clip-rule="evenodd"/><path d="M42 22V8a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMoveInOne0)"/>`),
		g.Group(children),
	)
}