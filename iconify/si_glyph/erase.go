package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Erase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.932 13.014L3.958 7.039L10.84.158c.322-.325.856-.314 1.191.022l4.762 4.759c.334.336.345.869.021 1.191l-6.882 6.884zm-.969 1.096c-1.582 1.583-5.434.3-7.102-1.368c-1.666-1.667-.52-3.087 1.063-4.67l6.039 6.038z"/>`),
		g.Group(children),
	)
}