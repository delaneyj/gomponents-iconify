package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21.725v-9.15L3 7.95v8.025q0 .55.263 1T4 17.7l7 4.025Zm2 0l7-4.025q.475-.275.738-.725t.262-1V7.95l-8 4.625v9.15Zm3.975-13.75l2.95-1.725L13 2.275Q12.525 2 12 2t-1 .275L9.025 3.4l7.95 4.575ZM12 10.85l2.975-1.7l-7.925-4.6l-3 1.725L12 10.85Z"/>`),
		g.Group(children),
	)
}