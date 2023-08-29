package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landslide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.47 13.79l-2.58-1.03L6 15.05L2 13.5v2.11l4 1.34l9.47-3.16m-4.9-2.37L8 8H2v3.61l4 1.34l4.57-1.53M6 19.05l-4-1.33V22h20l-4.97-6.62L6 19.05M17 6V1l-5-1l-3 2v4l3 2l5-2m1.5 1L16 9v3l2.5 2l4.5-2V8l-4.5-1Z"/>`),
		g.Group(children),
	)
}