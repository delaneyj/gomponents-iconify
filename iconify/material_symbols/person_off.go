package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L17.15 20H4v-2.8q0-.85.438-1.563T5.6 14.55q1.125-.575 2.288-.925t2.362-.525L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM18.4 14.55q.725.35 1.15 1.062T20 17.15l-3.35-3.35q.45.175.888.35t.862.4Zm-4.2-3.2L8.65 5.8q.575-.85 1.45-1.325T12 4q1.65 0 2.825 1.175T16 8q0 1.025-.475 1.9T14.2 11.35Z"/>`),
		g.Group(children),
	)
}