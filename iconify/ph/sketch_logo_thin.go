package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SketchLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m243 101.37l-56-64a4 4 0 0 0-3-1.37H72a4 4 0 0 0-3 1.37l-56 64a4 4 0 0 0 .09 5.36l112 120a4 4 0 0 0 5.84 0l112-120a4 4 0 0 0 .07-5.36ZM77.29 108l39.07 97.66L25.2 108Zm92.8 0L128 213.23L85.91 108ZM88 100l40-53.33L168 100Zm90.71 8h52.09l-91.16 97.66Zm52.47-8H178l-42-56h46.18ZM73.82 44H120l-42 56H24.82Z"/>`),
		g.Group(children),
	)
}