package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9v1a3 3 0 0 0 3 3h1m8-5v5a3 3 0 0 1-.129.872m-2.014 2a3 3 0 0 1-.857.124h-1M10 21V10m0-4V5a2 2 0 1 1 4 0v5m0 4v7m-7 0h10M3 3l18 18"/>`),
		g.Group(children),
	)
}