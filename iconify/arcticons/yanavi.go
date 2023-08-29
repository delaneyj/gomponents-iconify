package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yanavi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 5.5l-37 15.83l19 2.22l18-18.05Zm0 0l-18 18.05l2.23 19L42.5 5.5Zm-37 15.83L20 28l6.68 14.5l-2.23-18.95L5.5 21.33Z"/>`),
		g.Group(children),
	)
}