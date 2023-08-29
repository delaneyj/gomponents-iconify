package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 56v58.77c0 89.61-75.82 119.34-91 124.39a15.53 15.53 0 0 1-10 0c-15.2-5.05-91-34.78-91-124.39V56a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}