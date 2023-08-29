package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#786DBC"/><path fill="#FFF" d="m15.931 10.771l-1.629-1.633a1.144 1.144 0 0 1 1.582-1.654l.038.038l3.636 3.636c.447.447.447 1.17 0 1.617l-3.64 3.636a1.143 1.143 0 0 1-1.616-1.616l1.707-1.707a5.695 5.695 0 1 0 4.705 5.63V6.142a1.143 1.143 0 0 1 2.286 0v12.729a8 8 0 1 1-7.07-8.104l.001.004z"/></g>`),
		g.Group(children),
	)
}