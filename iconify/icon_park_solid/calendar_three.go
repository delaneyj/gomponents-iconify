package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCalendarThree0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M4 20h40M4 32h40"/><path stroke="#fff" d="M17 4v8m14-8v8"/><path stroke="#000" d="M17 20v24m14-24v24"/><path stroke="#fff" d="M44 13v26M4 13v26m10 5h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCalendarThree0)"/>`),
		g.Group(children),
	)
}