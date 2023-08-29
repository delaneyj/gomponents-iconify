package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.569 3.886A3 3 0 0 1 6.354 2h11.292a3 3 0 0 1 2.785 1.886l.895 2.236a3 3 0 0 1-1.122 3.61L20 9.87V20h1a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2h1V9.869l-.204-.137a3 3 0 0 1-1.122-3.61l.895-2.236ZM6 10.596V20h3v-4a3 3 0 1 1 6 0v4h3v-9.404c-.58 0-1.16-.168-1.664-.504L15 9.202l-1.336.89a3 3 0 0 1-3.328 0L9 9.202l-1.336.89A2.997 2.997 0 0 1 6 10.596ZM13 20v-4a1 1 0 1 0-2 0v4h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}