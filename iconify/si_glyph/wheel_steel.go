package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelSteel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.999 16C3.589 16 0 12.41 0 8s3.588-8 7.999-8C12.411 0 16 3.59 16 8s-3.59 8-8.001 8zM8 2C4.69 2 2 4.692 2 8s2.692 6 6 6s6-2.692 6-6s-2.69-6-6-6z"/><path d="M7.992 6c-2.316 0-4.098.797-4.961 2.346c.037.473.154.928.336 1.352c1.221-.652 3.551 1.83 2.818 2.935c.58.204 1.141.355 1.797.355a5.94 5.94 0 0 0 1.887-.315c-.734-1.105 1.525-3.535 2.754-2.869a4.31 4.31 0 0 0 .344-1.396C12.095 6.875 10.293 6 7.992 6zm.01 3.156c-.625 0-1.127-.51-1.127-1.141s.502-1.141 1.127-1.141c.619 0 1.123.51 1.123 1.141s-.504 1.141-1.123 1.141z"/></g>`),
		g.Group(children),
	)
}