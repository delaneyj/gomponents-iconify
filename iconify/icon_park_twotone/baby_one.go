package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBabyOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><circle cx="24" cy="8" r="5" fill="#555" stroke-linejoin="round"/><path d="M5 28s17-20.25 38 0"/><path fill="#555" stroke-linejoin="round" d="M19 28v-3.79S19 19 24 19s5 5.21 5 5.21V32s0 5-5 5s-5-5-5-5v-4Z"/><path stroke-linejoin="round" d="m29 32l8 5l-6 7M19 32l-8 5l6 7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBabyOne0)"/>`),
		g.Group(children),
	)
}