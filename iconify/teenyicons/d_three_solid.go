package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DThreeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2H0V1h1.5a6.5 6.5 0 0 1 0 13H0v-1h1.5a5.5 5.5 0 1 0 0-11Zm10 0H7V1h4.5a3.5 3.5 0 0 1 1.804 6.5A3.5 3.5 0 0 1 11.5 14H7v-1h4.5a2.5 2.5 0 0 0 0-5H9V7h2.5a2.5 2.5 0 0 0 0-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}