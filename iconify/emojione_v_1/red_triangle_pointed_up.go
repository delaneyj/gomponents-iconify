package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M62.46 55.1C53.353 38.468 44.245 21.84 35.138 5.204a3.576 3.576 0 0 0-6.325 0C19.706 21.836 10.596 38.464 1.491 55.1c-1.409 2.569.387 5.899 3.16 5.899h54.653c1.537 0 2.569-.828 3.108-1.952c.679-1.067.859-2.487.055-3.951"/>`),
		g.Group(children),
	)
}