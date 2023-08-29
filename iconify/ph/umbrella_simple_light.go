package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 126.79A110.43 110.43 0 0 0 53.11 55.22a109.51 109.51 0 0 0-35.06 71.57A14 14 0 0 0 32 142h90v58a30 30 0 0 0 60 0a6 6 0 0 0-12 0a18 18 0 0 1-36 0v-58h90a14 14 0 0 0 14-15.21Zm-12.49 2.56a2 2 0 0 1-1.51.65H32a2 2 0 0 1-1.49-.65a2 2 0 0 1-.53-1.56a98.43 98.43 0 0 1 164.78-63.74A97.5 97.5 0 0 1 226 127.79a2 2 0 0 1-.54 1.56Z"/>`),
		g.Group(children),
	)
}