package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M6 8.256L24.009 3L42 8.256v10.778A26.316 26.316 0 0 1 24.003 44A26.32 26.32 0 0 1 6 19.029V8.256Z"/>`),
		g.Group(children),
	)
}