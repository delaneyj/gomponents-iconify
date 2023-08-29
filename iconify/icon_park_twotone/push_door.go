package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPushDoor0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6h36v36H6"/><path fill="#555" d="M6 6v36l18-6V12L6 6Z"/><path d="M18 22v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPushDoor0)"/>`),
		g.Group(children),
	)
}