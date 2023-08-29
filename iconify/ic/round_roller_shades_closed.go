package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundRollerShadesClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19V5c0-1.1-.9-2-2-2H6c-1.1 0-2 .9-2 2v14H3c-.55 0-1 .45-1 1s.45 1 1 1h7.25c0 .97.78 1.75 1.75 1.75s1.75-.78 1.75-1.75H21c.55 0 1-.45 1-1s-.45-1-1-1h-1zM6 19v-2h5v2H6zm7 0v-2h5v2h-5z"/>`),
		g.Group(children),
	)
}