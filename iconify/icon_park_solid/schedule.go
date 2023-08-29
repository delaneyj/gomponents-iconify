package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSchedule0"><g fill="none" stroke-linecap="round" stroke-width="4"><rect width="40" height="30" x="4" y="10" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="2"/><path stroke="#fff" d="M14 6v8"/><path stroke="#000" d="M25 23H14m20 8H14"/><path stroke="#fff" d="M34 6v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSchedule0)"/>`),
		g.Group(children),
	)
}