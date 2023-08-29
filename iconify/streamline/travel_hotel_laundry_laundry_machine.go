package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelLaundryLaundryMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="7" cy="8" r="3.5"/><path d="M.5 3.5h3m7 0h3M7 7.97v3.5M4 6.2l3 1.77m3-1.77L7 7.97"/></g>`),
		g.Group(children),
	)
}