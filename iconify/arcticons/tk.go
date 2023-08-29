package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 8.75h23.72v7.78H20.9v11.34h-9.08V16.53H4.5Zm27.78 7.75H43.5L32.12 27.87L43.5 39.25H32.28L20.9 27.87Zm0 0"/>`),
		g.Group(children),
	)
}