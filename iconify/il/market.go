package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Market(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M648 256q0 2 1 3t0 3v20q0 10-7 17t-17 7h-46v324q0 10-6 16t-17 7H93q-10 0-17-7t-7-16V306H23q-10 0-16-7t-7-17v-20q0-4 1-6L60 39q5-16 17-25t28-9h439q16 0 28 9t16 25zm-138 50H139v127q0 5 4 8t8 4h347q5 0 9-4t3-8V306z"/>`),
		g.Group(children),
	)
}