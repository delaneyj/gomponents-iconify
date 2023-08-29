package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.363 3.863l20.5 20.5l-20.5 20.499l-20.5-20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.53 19.03l-7.417 7.416l-7.75 18.417"/>`),
		g.Group(children),
	)
}