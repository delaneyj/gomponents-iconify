package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMetricsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22v-5H2v-4h6.45l1.7 2.575h1.8l1.35-4.325L14.45 13H22v4h-5v5H7Zm3.7-9.25L9.525 11H2V7h5V2h10v5h5v4h-6.475l-1.7-2.55H12.05l-1.35 4.3Z"/>`),
		g.Group(children),
	)
}