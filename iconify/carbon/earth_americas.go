package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthAmericas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 16a14.031 14.031 0 1 0-7.14 12.191l.117.026l.024-.111A13.998 13.998 0 0 0 30 16ZM4 16a11.937 11.937 0 0 1 .395-3h4.243l4.992 4.16l-1.91 2.546a2.009 2.009 0 0 0 .186 2.614L14 24.414v3.405A12.01 12.01 0 0 1 4 16Zm9.32 4.906l3.05-4.066L9.362 11H5.106a11.962 11.962 0 0 1 17.778-4.814L22.279 8h-5.693l-3.043 3.043l9.354 8.313l-1.649 7.419A11.903 11.903 0 0 1 16 28v-4.414Zm10.315 4.344l1.213-5.46a2 2 0 0 0-.623-1.927l-7.768-6.906l.957-.957h4.865a1.998 1.998 0 0 0 1.898-1.368l.353-1.06a11.978 11.978 0 0 1-.895 17.678Z"/>`),
		g.Group(children),
	)
}