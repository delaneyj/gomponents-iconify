package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.097 0h10.025v24zm-7.063 0H0v24zm-.853 19.171l4.384-10.329l6.386 15.156h-4.184l-1.91-4.827z"/>`),
		g.Group(children),
	)
}