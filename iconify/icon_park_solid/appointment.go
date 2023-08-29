package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appointment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAppointment0"><g fill="none" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 41c0-8.837 8.059-16 18-16"/><circle cx="34" cy="34" r="9" fill="#fff" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M33 31v4h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAppointment0)"/>`),
		g.Group(children),
	)
}