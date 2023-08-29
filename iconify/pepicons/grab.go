package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 8a1 1 0 0 1 0-2h15a1 1 0 1 1 0 2h-15Zm0 3.25a1 1 0 1 1 0-2h15a1 1 0 1 1 0 2h-15Zm0 3.25a1 1 0 1 1 0-2h15a1 1 0 1 1 0 2h-15Z"/>`),
		g.Group(children),
	)
}