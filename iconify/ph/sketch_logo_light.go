package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SketchLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m244.52 100.05l-56-64A6 6 0 0 0 184 34H72a6 6 0 0 0-4.52 2l-56 64a6 6 0 0 0 .13 8l112 120a6 6 0 0 0 8.78 0l112-120a6 6 0 0 0 .13-7.95ZM75.94 110l34.6 86.49L29.81 110Zm91.2 0L128 207.84L88.86 110ZM92 98l36-48l36 48Zm88.06 12h46.13l-80.73 86.49Zm46.72-12H179l-39-52h41.28ZM74.72 46H116L77 98H29.22Z"/>`),
		g.Group(children),
	)
}