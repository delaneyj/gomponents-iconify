package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Badminton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBadminton0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M19.5 44c2.49 0 4.5-2.01 4.5-4.5V19h-4l-5 20.5c0 2.49 2.01 4.5 4.5 4.5Z"/><path d="M20 19h-4L6.68 37.9c-1.11 2.61.48 5.74 3.3 6.07A4.494 4.494 0 0 0 15 39.5M28.5 44c-2.49 0-4.5-2.01-4.5-4.5V19h4l5 20.5c0 2.49-2.01 4.5-4.5 4.5Z"/><path d="M28 19h4l9.32 18.9c1.11 2.61-.48 5.74-3.3 6.07A4.494 4.494 0 0 1 33 39.5"/><path fill="#fff" d="M16.06 13c-.04-.33-.06-.66-.06-1c0-4.42 3.58-8 8-8s8 3.58 8 8c0 .34-.02.67-.06 1H16.06Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBadminton0)"/>`),
		g.Group(children),
	)
}