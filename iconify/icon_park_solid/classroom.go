package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Classroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClassroom0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="13" r="9" fill="#fff"/><path d="M5 44c0-8.438 6.175-16.313 11.4-18c0 0 4.75 5.063 7.6 8.438L31.6 26c4.275.563 11.4 9.563 11.4 18"/><path stroke-linecap="round" d="M2 44h44"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClassroom0)"/>`),
		g.Group(children),
	)
}