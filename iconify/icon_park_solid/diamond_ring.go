package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDiamondRing0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="25" cy="29" r="15"/><path fill="#fff" d="m18 8l3-4h8.054L32 8l-7 6l-7-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDiamondRing0)"/>`),
		g.Group(children),
	)
}