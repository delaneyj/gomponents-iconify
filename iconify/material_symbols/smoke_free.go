package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokeFree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L13.2 16H2v-3h8.2L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Zm-.95-6.6l-.85-.85V13h1.5v3h-.65Zm1.65 0v-3H22v3h-1.5ZM17 14.15L15.85 13H17v1.15ZM18 12v-1.3q0-.975-.575-1.488T16.05 8.7H14.5q-1.4 0-2.375-.975T11.15 5.35q0-1.4.975-2.375T14.5 2v1.5q-.75 0-1.3.525t-.55 1.325q0 .8.55 1.325t1.3.525h1.55q1.4 0 2.425.9t1.025 2.25V12H18Zm2.5 0V9.75q0-1.65-1.15-2.85T16.5 5.7V4.2q.75 0 1.3-.55t.55-1.3h1.5q0 .725-.275 1.313T18.85 4.75q1.4.65 2.275 2t.875 3V12h-1.5Z"/>`),
		g.Group(children),
	)
}