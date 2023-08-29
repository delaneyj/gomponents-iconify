package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#D3D3D3" d="M16 30.064c7.732 0 14-3.506 14-7.974S23.732 14 16 14S2 17.622 2 22.09s6.268 7.974 14 7.974Z"/><path fill="#321B41" d="M16 30.11c7.185 0 13.01-3.138 13.01-7.01c0-3.871-5.825-7.01-13.01-7.01S2.99 19.23 2.99 23.1c0 3.872 5.825 7.01 13.01 7.01Z"/></g>`),
		g.Group(children),
	)
}