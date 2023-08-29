package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VacuumCleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVacuumCleaner0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M26 22.5V10c0-3 2-6 6-6s6 3 6 6v24"/><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="M32.75 34h10.5l.75 6H32l.75-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 40h15.886c.063 0 .114-.05.114-.114V23.255C26 15.935 20.066 10 12.745 10v0A5.745 5.745 0 0 0 7 15.745V29"/><circle cx="10" cy="34" r="6" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 10v19"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVacuumCleaner0)"/>`),
		g.Group(children),
	)
}