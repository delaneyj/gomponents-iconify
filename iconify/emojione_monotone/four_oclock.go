package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm14.067 41L34 35.535V38h-4v-2.555c-1.19-.693-2-1.969-2-3.445c0-.057.015-.11.017-.166L26 30.586L27.93 27L30 28.281V6h4v22.555a3.983 3.983 0 0 1 1.999 3.438L48 39.416L46.067 43z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}