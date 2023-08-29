package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 1)"><path d="m13.842 4l-3.93-2.85a1.495 1.495 0 0 0-2.928.093L2.918 4H.021v9.96h15.938V4h-2.117zM8.458 3c.556 0 1.034-.305 1.294-.753L11.895 4H5.047l2.166-1.664c.27.4.727.664 1.245.664zm3.503 7.085L8.72 11.758l-3.727-4.72L1 13V5h14v8l-3.04-2.915z"/><circle cx="12.963" cy="6.963" r=".963"/></g>`),
		g.Group(children),
	)
}