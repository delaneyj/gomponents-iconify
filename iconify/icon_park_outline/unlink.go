package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m25.893 16.03l-7.779-7.778c-2.863-2.863-7.41-2.959-10.157-.213c-2.746 2.746-2.65 7.293.214 10.157l7.778 7.778m15.967-3.904l7.778 7.779c2.864 2.864 3.235 7.274.214 10.157s-7.294 2.65-10.157-.213l-7.779-7.779m-.734-10.938l-3.889-3.89M30.313 30.15l-3.889-3.889"/>`),
		g.Group(children),
	)
}