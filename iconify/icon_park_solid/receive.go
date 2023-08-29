package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSReceive0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4.592 19.467A2 2 0 0 1 6.537 17h34.926a2 2 0 0 1 1.945 2.467l-5.04 21A2 2 0 0 1 36.423 42H11.577a2 2 0 0 1-1.945-1.533l-5.04-21Z"/><path stroke="#fff" stroke-linejoin="round" d="M11 7h8v10h-8zm8 10l6.5-9L38 17"/><path stroke="#000" d="M15 25h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSReceive0)"/>`),
		g.Group(children),
	)
}