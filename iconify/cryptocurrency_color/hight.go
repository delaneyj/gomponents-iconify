package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#117FC0"/><path fill="#FFF" d="m26.496 10.332l-8.455 13.94h-5.042l3.875-6.397h-3.851l-4.49 7.404a11.905 11.905 0 0 0 17.96-14.94l.003-.007zM19.114 7.83l-3.67 6.04H19.3l4.274-7.049a11.905 11.905 0 0 0-17.96 14.94l8.452-13.93h5.047z"/></g>`),
		g.Group(children),
	)
}