package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M27 67Q5 67 5 88v256q0 21 22 21q21 0 21-21v-73l134 150q8 6 15 6q5 0 9-2q13-3 13-19V24q0-15-13-19t-24 6L48 161V88q0-21-21-21zm149 12v271L54 216z"/>`),
		g.Group(children),
	)
}