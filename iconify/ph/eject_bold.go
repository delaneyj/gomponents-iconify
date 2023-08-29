package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 160H48a20 20 0 0 0-20 20v20a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20v-20a20 20 0 0 0-20-20Zm-4 36H52v-12h152ZM48.27 144h159.46a20.27 20.27 0 0 0 14.61-34.3l-73.76-76.92a28.51 28.51 0 0 0-41.16 0L33.66 109.7A20.27 20.27 0 0 0 48.27 144Zm76.48-94.61a4.49 4.49 0 0 1 6.5 0L199 120H57Z"/>`),
		g.Group(children),
	)
}