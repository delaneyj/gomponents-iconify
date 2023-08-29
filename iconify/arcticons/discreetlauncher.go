package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discreetlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.13 22.25L24 4.5L6.87 22.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.13 30.54L24 12.79L6.87 30.54"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.73 30.37V43.5h20.54V30.37"/>`),
		g.Group(children),
	)
}