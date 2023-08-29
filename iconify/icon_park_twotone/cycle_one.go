package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCycleOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 20c0-8 4-12 12-12m22 22c0 8-4 12-12 12"/><path fill="#555" d="M28 18c0-5.523 4.477-10 10-10h4v14H28v-4ZM6 28h14v4c0 5.523-4.477 10-10 10H6V28Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCycleOne0)"/>`),
		g.Group(children),
	)
}