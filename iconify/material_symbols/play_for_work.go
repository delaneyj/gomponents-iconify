package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayForWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-2.5 0-4.25-1.75T6 14h2q0 1.65 1.175 2.825T12 18q1.65 0 2.825-1.175T16 14h2q0 2.5-1.75 4.25T12 20Zm0-5.025l-4-4L9.4 9.55l1.6 1.6V5h2v6.15l1.6-1.6l1.4 1.425l-4 4Z"/>`),
		g.Group(children),
	)
}