package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 .062c-4.419 0-8 3.559-8 7.947c0 4.39 3.581 7.949 8 7.949c4.418 0 8-3.56 8-7.949C16 3.621 12.418.062 8 .062zm3.108 11.963L8.021 9.902l-3.088 2.123L6.112 8.59L3.024 6.465h3.817l1.18-3.435l1.18 3.435h3.816L9.93 8.59l1.178 3.435z"/>`),
		g.Group(children),
	)
}