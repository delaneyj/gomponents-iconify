package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSchedule0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="40" height="30" x="4" y="10" fill="#555" stroke-linejoin="round" rx="2"/><path d="M14 6v8m11 9H14m20 8H14M34 6v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSchedule0)"/>`),
		g.Group(children),
	)
}