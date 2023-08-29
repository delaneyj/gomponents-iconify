package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sokoban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l-8.91 25.73h17.82Zm-8.91 25.73L24 43.5l8.91-13.27Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l8.91 25.73l8.87-.39Zm17.78 25.34l-8.87.39L24 43.5ZM24 4.5L6.22 29.84l8.87.39Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.22 29.84L24 43.5l-8.91-13.27Z"/>`),
		g.Group(children),
	)
}