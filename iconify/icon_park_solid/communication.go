package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Communication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCommunication0"><g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#fff" stroke-linejoin="round" d="M33 38H22v-8h14v-8h8v16h-5l-3 3l-3-3Z"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4 6h32v24H17l-4 4l-4-4H4V6Z"/><path stroke="#000" d="M19 18h1m6 0h1m-15 0h1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCommunication0)"/>`),
		g.Group(children),
	)
}