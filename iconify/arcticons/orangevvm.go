package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orangevvm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 16a8.31 8.31 0 0 0-8 8a8.3 8.3 0 0 0 8 8a8.3 8.3 0 0 0 8-8a8.3 8.3 0 0 0-8-8Zm24 0a8.3 8.3 0 0 0-8 8a8.3 8.3 0 0 0 8 8a8.31 8.31 0 0 0 8-8a8.3 8.3 0 0 0-8-8ZM11.99 31.99h24.02"/>`),
		g.Group(children),
	)
}