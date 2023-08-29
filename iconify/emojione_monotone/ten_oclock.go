package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4.068 35L34 35.721V38h-4v-2.555a3.983 3.983 0 0 1-1.999-3.437L16 24.584L17.932 21L30 28.467V6h4v22.555A3.98 3.98 0 0 1 36 32c0 .058-.015.111-.017.168L38 33.416L36.068 37z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}