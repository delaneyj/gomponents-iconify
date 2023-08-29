package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 213h150v150q0 21 21 21t21-21V213h150q21 0 21-21t-21-21H213V21q0-21-21-21t-21 21v150H21q-21 0-21 21t21 21z"/>`),
		g.Group(children),
	)
}