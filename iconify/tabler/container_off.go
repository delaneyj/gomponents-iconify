package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4v.01M20 20v.01M20 16v.01M20 12v.01M20 8v.01M8.297 4.289A1 1 0 0 1 9 4h6a1 1 0 0 1 1 1v7m0 4v3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V8M4 4v.01M4 20v.01M4 16v.01M4 12v.01M4 8v.01M3 3l18 18"/>`),
		g.Group(children),
	)
}