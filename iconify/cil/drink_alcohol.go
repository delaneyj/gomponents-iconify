package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrinkAlcohol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 440h-72v32h176v-32h-72V318.968l176-192.762V80H72v46.206l176 192.762ZM104 113.794V112h320v1.794L374.508 168H153.492ZM182.709 200h162.582l-80.349 88h-1.884Z"/>`),
		g.Group(children),
	)
}