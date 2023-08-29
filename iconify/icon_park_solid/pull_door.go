package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PullDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPullDoor0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 8v32l18.2 4V4L6 8Z"/><path stroke="#fff" d="M24.2 8H42v32H24.2"/><path stroke="#000" d="M18 22v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPullDoor0)"/>`),
		g.Group(children),
	)
}