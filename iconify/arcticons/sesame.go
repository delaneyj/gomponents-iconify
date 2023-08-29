package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sesame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.47 4.5a12.138 12.138 0 0 0 0 24.277h7.196V4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.334 16.638V43.5h7.196a12.137 12.137 0 0 0 12.136-12.138V16.638"/>`),
		g.Group(children),
	)
}