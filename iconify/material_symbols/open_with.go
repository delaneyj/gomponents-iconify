package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenWith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-4.25-4.25l1.425-1.425L11 18.15V14h2l-.025 4.125L14.8 16.3l1.45 1.45L12 22Zm-5.75-5.75L2 12l4.225-4.225L7.65 9.2L5.85 11H10v2H5.875L7.7 14.8l-1.45 1.45Zm11.5 0l-1.425-1.425L18.15 13H14v-2l4.125.025L16.3 9.2l1.45-1.45L22 12l-4.25 4.25ZM11 10V5.85L9.175 7.675L7.75 6.25L12 2l4.25 4.25l-1.425 1.425L13 5.85V10h-2Z"/>`),
		g.Group(children),
	)
}