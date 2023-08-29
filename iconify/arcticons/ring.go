package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.68 3.51a1.59 1.59 0 0 0-1 .23L9.62 8.4A1.76 1.76 0 0 0 9 10.8l3.88 6.79h-7.8a1.75 1.75 0 0 0-1.76 1.75v9.31a1.75 1.75 0 0 0 1.76 1.75h7.78L9 37.19a1.78 1.78 0 0 0 .64 2.42l8.09 4.66a1.76 1.76 0 0 0 2.42-.68L24 36.8l3.88 6.79a1.78 1.78 0 0 0 2.43.68l8-4.66a1.76 1.76 0 0 0 .68-2.42l-3.82-6.79h7.76a1.75 1.75 0 0 0 1.75-1.75v-9.31a1.75 1.75 0 0 0-1.75-1.75h-7.76l3.88-6.79a1.75 1.75 0 0 0-.68-2.4l-8-4.66a1.78 1.78 0 0 0-2.43.64L24 11.19l-3.87-6.81a1.78 1.78 0 0 0-1.45-.87Z"/>`),
		g.Group(children),
	)
}