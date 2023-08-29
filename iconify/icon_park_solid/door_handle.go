package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorHandle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDoorHandle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="26" height="40" x="6" y="4" rx="2"/><path d="M14 34h10"/><path fill="#fff" d="M42 20v-6H23a5 5 0 1 0 0 6h19Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDoorHandle0)"/>`),
		g.Group(children),
	)
}