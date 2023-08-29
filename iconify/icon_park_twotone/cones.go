package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCones0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 8L4 40h40L24 8Z"/><path stroke-linecap="round" d="m30 32l-6-4l-6 4m6-4v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCones0)"/>`),
		g.Group(children),
	)
}