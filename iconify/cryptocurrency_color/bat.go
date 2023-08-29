package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#FF5000"/><path fill="#FFF" d="m6 23.5l10.051-17L26 23.477L6 23.5zm10.027-10.12l-4.108 6.786h8.235l-4.127-6.786z"/></g>`),
		g.Group(children),
	)
}