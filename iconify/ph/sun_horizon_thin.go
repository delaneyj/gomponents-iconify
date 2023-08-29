package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunHorizonThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 156h-45.06A68 68 0 1 0 60 144a68.73 68.73 0 0 0 1.06 12H16a4 4 0 0 0 0 8h224a4 4 0 0 0 0-8ZM68 144a60 60 0 1 1 118.79 12H69.21A60.16 60.16 0 0 1 68 144Zm144 56a4 4 0 0 1-4 4H48a4 4 0 0 1 0-8h160a4 4 0 0 1 4 4ZM76.42 41.79a4 4 0 0 1 7.16-3.58l8 16a4 4 0 0 1-7.16 3.58Zm-56 52.42a4 4 0 0 1 5.37-1.79l16 8a4 4 0 0 1-3.58 7.16l-16-8a4 4 0 0 1-1.79-5.37Zm192 11.58a4 4 0 0 1 1.79-5.37l16-8a4 4 0 1 1 3.58 7.16l-16 8a4 4 0 0 1-5.37-1.79Zm-48-51.58l8-16a4 4 0 1 1 7.16 3.58l-8 16a4 4 0 0 1-7.16-3.58Z"/>`),
		g.Group(children),
	)
}