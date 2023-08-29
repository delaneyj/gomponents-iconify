package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 45h64v171H192v192q0 9-6.5 15t-14.5 6h-43q-9 0-15-6t-6-15V173h213V88h-21v21q0 9-6.5 15.5T277 131H21q-8 0-14.5-6.5T0 109V24q0-9 6.5-15T21 3h256q9 0 15.5 6t6.5 15v21z"/>`),
		g.Group(children),
	)
}