package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libretube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.773 24l-8.661-5v10l8.661-5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.112 4.5v9L25.299 24L7.112 34.5v9L40.888 24L7.112 4.5z"/>`),
		g.Group(children),
	)
}