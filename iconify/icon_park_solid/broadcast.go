package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Broadcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBroadcast0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M24 28.63a4.538 4.538 0 0 0 4.546-4.532A4.538 4.538 0 0 0 24 19.567a4.538 4.538 0 0 0-4.545 4.53A4.538 4.538 0 0 0 24 28.63Z"/><path stroke-linecap="round" d="M16 15c-5.333 4.97-5.333 13.03 0 18m16 0c5.333-4.97 5.333-13.03 0-18M9.858 10c-7.81 7.786-7.81 20.41 0 28.196m28.284 0c7.81-7.786 7.81-20.41 0-28.196"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBroadcast0)"/>`),
		g.Group(children),
	)
}