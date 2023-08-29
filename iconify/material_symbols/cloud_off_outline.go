package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L17.15 20H6.5q-2.3 0-3.9-1.6T1 14.5q0-1.925 1.188-3.425T5.25 9.15q.075-.2.15-.388t.15-.412L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM6.5 18h8.65L7.1 9.95q-.05.275-.075.525T7 11h-.5q-1.45 0-2.475 1.025T3 14.5q0 1.45 1.025 2.475T6.5 18Zm4.625-4.025ZM21.6 18.75l-1.45-1.4q.425-.35.638-.813T21 15.5q0-1.05-.725-1.775T18.5 13H17v-2q0-2.075-1.463-3.538T12 6q-.675 0-1.3.163t-1.2.512l-1.45-1.45q.875-.6 1.863-.912T12 4q2.925 0 4.963 2.038T19 11q1.725.2 2.863 1.488T23 15.5q0 .975-.375 1.813T21.6 18.75Zm-6.775-6.725Z"/>`),
		g.Group(children),
	)
}