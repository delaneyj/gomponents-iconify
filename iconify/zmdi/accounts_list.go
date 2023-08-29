package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountsList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 43h42v298h-42V43zm-86 298V43h43v298h-43zM277 43q9 0 15.5 6t6.5 15v256q0 9-6.5 15t-15.5 6H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43h256zm-128 58q-20 0-34 14t-14 34t14 34t34 14t34-14t14-34t-14-34t-34-14zm96 198v-16q0-22-33-35t-63-13t-63 13t-33 35v16h192z"/>`),
		g.Group(children),
	)
}