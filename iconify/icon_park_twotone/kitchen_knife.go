package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTKitchenKnife0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 30v10c0 6 8 6 8 0V30m0 7h-8"/><path fill="#555" stroke="#fff" stroke-width="4" d="M14 6a2 2 0 0 1 2-2h16.635c.319 0 .632.075.888.265C34.542 5.025 37.198 7.582 38 14c.773 6.182-1.369 12.364-2.382 14.855c-.288.71-.985 1.145-1.75 1.145H14V6Z"/><circle cx="22" cy="10" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTKitchenKnife0)"/>`),
		g.Group(children),
	)
}