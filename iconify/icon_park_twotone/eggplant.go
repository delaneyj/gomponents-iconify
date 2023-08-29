package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eggplant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEggplant0"><g fill="none" stroke="#fff" stroke-width="4"><path d="m35 18l-3.88 12.612c-.08.259-.155.514-.248.768c-.548 1.498-2.99 7.04-9.871 9.62C13 44 4.999 40 5 32.077C5 24.154 13 25 17 22s8-9 8-9"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M39.255 24.488S40.371 19.176 40 16c-.505-4.327-2.5-7.5-5-9c-2.502-1.5-7.748-2.5-12 0c-4.253 2.5-6.033 5.845-6.033 5.845l8.206-.767a1 1 0 0 1 1.055.72l1.497 5.24a1 1 0 0 0 1.236.688l5.282-1.51a1 1 0 0 1 1.145.468l3.865 6.804Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m36 8l6-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEggplant0)"/>`),
		g.Group(children),
	)
}