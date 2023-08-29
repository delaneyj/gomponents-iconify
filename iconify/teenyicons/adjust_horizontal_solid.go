package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustHorizontalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 0H6v2H0v1h6v2h3V3h6V2H9V0ZM5 5H2v2H0v1h2v2h3V8h10V7H5V5Zm8 5h-3v2H0v1h10v2h3v-2h2v-1h-2v-2Z"/>`),
		g.Group(children),
	)
}