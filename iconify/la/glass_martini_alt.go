package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassMartiniAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.5 4A5.503 5.503 0 0 0 3 9.5C3 12.535 5.465 15 8.5 15c.813 0 1.578-.203 2.281-.531L16 21.344V27h-4v2h10v-2h-4v-5.656l8.813-11.625l.187-.281V8H13.75c-.605-2.32-2.766-4-5.25-4zm0 2c1.39 0 2.59.84 3.125 2H7v1.438l.188.28l2.343 3.095A3.131 3.131 0 0 1 8.5 13A3.46 3.46 0 0 1 5 9.5A3.46 3.46 0 0 1 8.5 6zm1.438 4h14.124l-1.5 2H13.5l1.5 2h6.031L17 19.344z"/>`),
		g.Group(children),
	)
}