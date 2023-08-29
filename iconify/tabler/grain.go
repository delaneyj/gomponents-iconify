package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.5 9.5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m5-5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m0 10a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-5 5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m10-10a1 1 0 1 0 2 0a1 1 0 1 0-2 0m5-5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-5 15a1 1 0 1 0 2 0a1 1 0 1 0-2 0m5-5a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/>`),
		g.Group(children),
	)
}