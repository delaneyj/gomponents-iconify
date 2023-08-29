package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstagramOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInstagramOne0"><g fill="#555" stroke="#fff" stroke-width="4"><circle cx="9" cy="8" r="4"/><path stroke-linejoin="round" d="M5 18h8v25H5zm16 9.5V43h7V29c0-2.5 1.5-4.5 4-4.5s4 2.5 4 4.5v14h7V27.5c0-3-3.5-9.5-11-9.5s-11 6.5-11 9.5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInstagramOne0)"/>`),
		g.Group(children),
	)
}