package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.619L4.5 24.964h5.85V40.38h27.3V24.964h5.85Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.901 28.817a3.9 3.9 0 0 1-7.8 0c0-2.12 3.9-7.708 3.9-7.708s3.9 5.589 3.9 7.708Z"/>`),
		g.Group(children),
	)
}