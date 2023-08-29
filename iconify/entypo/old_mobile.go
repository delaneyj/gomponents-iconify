package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OldMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.6 3H7V0H5v18.6c0 .77.629 1.4 1.398 1.4H13.6c.769 0 1.4-.631 1.4-1.4V4.401C15 3.629 14.369 3 13.6 3zM8 15c-.691 0-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1s-.559 1-1.25 1zm1.25 2c0 .553-.559 1-1.25 1s-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1zM7 11V5h6v6H7zm5 4c-.691 0-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1s-.559 1-1.25 1zm1.25 2c0 .553-.559 1-1.25 1s-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1z"/>`),
		g.Group(children),
	)
}