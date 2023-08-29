package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v240h1200V0H0zm0 480v240h1200V480H0zm0 480v240h1200V960H0z"/>`),
		g.Group(children),
	)
}