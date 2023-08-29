package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Android(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAndroid0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M43.901 36H4.1C5.103 25.893 13.63 18 24 18c10.372 0 18.899 7.893 19.902 18Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m14 20l-4-7m23 7l4-7"/><circle cx="15" cy="29" r="2" fill="#fff"/><circle cx="33" cy="29" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAndroid0)"/>`),
		g.Group(children),
	)
}