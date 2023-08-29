package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#ff1f43"/><path fill="#fff" d="M7 7h4.2l7.2 12.775V7H22v18h-4.8z"/></g>`),
		g.Group(children),
	)
}