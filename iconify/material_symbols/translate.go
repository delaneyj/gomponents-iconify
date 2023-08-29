package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.9 22l4.55-12h2.1l4.55 12H21l-1.05-3.05H15.1L14 22h-2.1Zm3.8-4.8h3.6l-1.75-4.95h-.1L15.7 17.2ZM4 19l-1.4-1.4l5.05-5.05q-.95-1.05-1.663-2.175T4.75 8h2.1q.45.9.963 1.625T9.05 11.15q1.1-1.2 1.825-2.463T12.1 6H1V4h7V2h2v2h7v2h-2.9q-.525 1.775-1.425 3.45T10.45 12.6l2.4 2.45l-.75 2.05L9 14l-5 5Z"/>`),
		g.Group(children),
	)
}