package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 10h14a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2H12a2.002 2.002 0 0 0-2 2v1H6V2H4v23a2.002 2.002 0 0 0 2 2h4v1a2.002 2.002 0 0 0 2 2h14a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2H12a2.002 2.002 0 0 0-2 2v1H6v-8h4v1a2.002 2.002 0 0 0 2 2h14a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2H12a2.002 2.002 0 0 0-2 2v1H6V7h4v1a2.002 2.002 0 0 0 2 2Zm0-6h14l.001 4H12Zm0 20h14l.001 4H12Zm0-10h14l.001 4H12Z"/>`),
		g.Group(children),
	)
}