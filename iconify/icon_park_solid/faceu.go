package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faceu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFaceu0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-width="4" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#000" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 12v4c2.667.833 9 3 9 10s-6 10-12 10s-12-3-12-10c0-8 9.667-12.833 15-14Z"/><ellipse cx="29" cy="26" fill="#000" rx="3" ry="4"/><ellipse cx="19" cy="26" fill="#000" rx="3" ry="4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFaceu0)"/>`),
		g.Group(children),
	)
}