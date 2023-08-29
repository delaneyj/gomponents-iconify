package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoPhotographySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-2.3-2.3H2V5h3.025v2.85L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM22 19.125l-5.525-5.55q.125-1.05-.213-2.037T15.176 9.8q-.725-.725-1.713-1.063t-2.037-.212L7.5 4.625L9 3h6l1.85 2H22v14.125ZM12 17.5q.575 0 1.113-.125t1.037-.4L8.025 10.85q-.275.5-.4 1.038T7.5 13q0 1.875 1.313 3.188T12 17.5Zm0-2q-1.05 0-1.775-.725T9.5 13q0-.5.188-.963t.537-.812l3.55 3.55q-.35.35-.812.538T12 15.5Z"/>`),
		g.Group(children),
	)
}