package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 3.538a6.461 6.461 0 0 1-1.884.516a3.301 3.301 0 0 0 1.444-1.816a6.607 6.607 0 0 1-2.084.797a3.28 3.28 0 0 0-2.397-1.034a3.28 3.28 0 0 0-3.197 4.028a9.321 9.321 0 0 1-6.766-3.431a3.284 3.284 0 0 0 1.015 4.381A3.301 3.301 0 0 1 .643 6.57v.041A3.283 3.283 0 0 0 3.277 9.83a3.291 3.291 0 0 1-1.485.057a3.293 3.293 0 0 0 3.066 2.281a6.586 6.586 0 0 1-4.862 1.359a9.286 9.286 0 0 0 5.034 1.475c6.037 0 9.341-5.003 9.341-9.341c0-.144-.003-.284-.009-.425a6.59 6.59 0 0 0 1.637-1.697z"/>`),
		g.Group(children),
	)
}