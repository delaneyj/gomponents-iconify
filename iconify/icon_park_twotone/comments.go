package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTComments0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M33 38H22v-8h14v-8h8v16h-5l-3 3l-3-3Z"/><path fill="#555" stroke-linejoin="round" d="M4 6h32v24H17l-4 4l-4-4H4V6Z"/><path d="M12 22h6m-6-8h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTComments0)"/>`),
		g.Group(children),
	)
}