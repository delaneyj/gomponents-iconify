package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parallelogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSParallelogram0"><path fill="#fff" stroke="#fff" stroke-width="4" d="M41.28 8H15.47a2 2 0 0 0-1.909 1.403l-8.75 28A2 2 0 0 0 6.721 40H32.53a2 2 0 0 0 1.909-1.404l8.75-28A2 2 0 0 0 41.279 8Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSParallelogram0)"/>`),
		g.Group(children),
	)
}