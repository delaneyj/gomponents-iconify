package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 8 7.484V9h2V8H8V6.5A1.5 1.5 0 0 1 9.5 5h.5V4h-.5A2.5 2.5 0 0 0 7 6.5V8H5v1h2v5.984A7.5 7.5 0 0 1 0 7.5Z"/>`),
		g.Group(children),
	)
}