package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-1.25-.125-2.4-.613T6.4 20.1l1.4-1.45q.725.525 1.538.85t1.662.45V22Zm2 0v-2.05q2.6-.375 4.3-2.337T19 13.05q0-2.925-2.038-4.963T12 6.05h-.2l1.6 1.6l-1.4 1.4l-4-4l4-4l1.4 1.45l-1.55 1.55H12q1.875 0 3.513.713t2.85 1.925q1.212 1.212 1.925 2.85T21 13.05q0 3.425-2.275 5.963T13 22Zm-8.05-3.35q-.8-1.05-1.288-2.2t-.612-2.4H5.1q.125.85.45 1.663t.85 1.537l-1.45 1.4Zm-1.9-6.6q.15-1.275.625-2.45T4.95 7.45l1.45 1.4q-.525.725-.85 1.538T5.1 12.05H3.05Z"/>`),
		g.Group(children),
	)
}