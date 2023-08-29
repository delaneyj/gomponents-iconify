package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbersLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a6 6 0 0 1-6 6H104a6 6 0 0 1 0-12h112a6 6 0 0 1 6 6ZM104 70h112a6 6 0 0 0 0-12H104a6 6 0 0 0 0 12Zm112 116H104a6 6 0 0 0 0 12h112a6 6 0 0 0 0-12ZM42.68 53.37L50 49.71V104a6 6 0 0 0 12 0V40a6 6 0 0 0-8.68-5.37l-16 8a6 6 0 0 0 5.36 10.74ZM72 202H52l21.48-28.74A21.5 21.5 0 0 0 77.79 157A21.75 21.75 0 0 0 69 142.38a22.86 22.86 0 0 0-31.35 4.31a22.18 22.18 0 0 0-3.28 5.92a6 6 0 0 0 11.28 4.11a9.87 9.87 0 0 1 1.48-2.67a10.78 10.78 0 0 1 14.78-2a9.89 9.89 0 0 1 4 6.61a9.64 9.64 0 0 1-2 7.28l-.06.09l-28.65 38.38A6 6 0 0 0 40 214h32a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}