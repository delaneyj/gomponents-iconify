package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="m5.5 11.493l2 2.998l2-2.998h4a1 1 0 0 0 1-1V1.5a.999.999 0 0 0-1-.999h-12a1 1 0 0 0-1 1v8.993a1 1 0 0 0 1 1h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}