package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.05 21.825q-.275 0-.525-.125t-.375-.4L7.35 14q-1.775.075-3.062-1.025T3 10q0-1.275.738-2.3T5.6 6.25q.45-2.275 2.238-3.763T12 1q2.375 0 4.163 1.488T18.4 6.25q1.125.425 1.863 1.45T21 10q0 1.875-1.325 2.975T16.7 14l-3.775 7.3q-.125.275-.363.4t-.512.125Zm0-3.175l2.7-5.25q-.6.3-1.3.45T12 14q-.675 0-1.363-.15T9.3 13.4l2.75 5.25Z"/>`),
		g.Group(children),
	)
}