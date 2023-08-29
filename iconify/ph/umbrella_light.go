package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 126.79A110.43 110.43 0 0 0 53.11 55.22a109.51 109.51 0 0 0-35.06 71.57A14 14 0 0 0 32 142h90v58a30 30 0 0 0 60 0a6 6 0 0 0-12 0a18 18 0 0 1-36 0v-58h90a14 14 0 0 0 14-15.21ZM94.11 130c1.69-51.21 24.7-80.16 33.89-89.73c9.2 9.58 32.2 38.53 33.89 89.73Zm-63.57-.65a2 2 0 0 1-.53-1.56a98.14 98.14 0 0 1 82.91-88.62c-12 15-29.43 44.44-30.83 90.83H32a2 2 0 0 1-1.46-.65Zm194.92 0a2 2 0 0 1-1.46.65h-50.09c-1.4-46.39-18.81-75.87-30.83-90.83A98.14 98.14 0 0 1 226 127.79a2 2 0 0 1-.54 1.56Z"/>`),
		g.Group(children),
	)
}