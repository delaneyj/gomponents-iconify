package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecurringAppointment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 0C5.838 0 0 5.838 0 13c0 7.163 5.838 13 13 13c6.555 0 12-4.888 12.875-11.219a1.5 1.5 0 1 0-2.969-.406C22.233 19.245 18.071 23 13 23C7.46 23 3 18.541 3 13C3 7.46 7.46 3 13 3c3.236 0 6.109 1.535 7.938 3.906l-1.782 1.782c-.15.15-.189.391-.125.593c.06.202.229.332.438.375c2.697.53 5.807.296 5.937.281a.623.623 0 0 0 .344-.187c.09-.088.17-.18.188-.313c.015-.13.25-3.267-.282-5.968a.55.55 0 0 0-.375-.438c-.204-.062-.444-.025-.593.125l-1.625 1.625C20.678 1.87 17.047 0 13 0z"/>`),
		g.Group(children),
	)
}