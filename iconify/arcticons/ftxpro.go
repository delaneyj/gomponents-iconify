package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ftxpro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 28.875H24v9.75h-9.75zm0-19.5H43.5v9.75H14.25zm-9.75 9.75h9.75v9.75H4.5zm9.75 0h19.5v9.75h-19.5z"/>`),
		g.Group(children),
	)
}