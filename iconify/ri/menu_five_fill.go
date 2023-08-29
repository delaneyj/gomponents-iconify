package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 18v2H6v-2h12Zm3-7v2H3v-2h18Zm-3-7v2H6V4h12Z"/>`),
		g.Group(children),
	)
}