package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pennant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 44h4M12 6V4v2Zm0 16v22v-22Zm0 22H8h4Zm-4 0h8M12 6v16l28-8l-28-8Z"/>`),
		g.Group(children),
	)
}