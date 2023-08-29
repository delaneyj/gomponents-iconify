package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.025 21.95q-3.85-.375-6.425-3.225T2.025 12q0-3.875 2.575-6.725t6.425-3.225v3q-2.6.35-4.3 2.325T5.025 12q0 2.65 1.7 4.625t4.3 2.325v3Zm2 0v-3q2.35-.3 3.975-1.95t1.975-4h3q-.35 3.575-2.863 6.088t-6.087 2.862ZM18.975 11Q18.625 8.65 17 7t-3.975-1.95v-3q3.575.35 6.088 2.863T21.974 11h-3Z"/>`),
		g.Group(children),
	)
}