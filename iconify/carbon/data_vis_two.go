package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataVisTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H17a2.002 2.002 0 0 0-2 2v6H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h11a2.002 2.002 0 0 0 2-2v-6h11a2.003 2.003 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2Zm0 2v4H17V4ZM15 22H4v-4h11Zm2-12h11l.001 4H17Zm-2 2v4H4v-4ZM4 28v-4h11.001v4Zm13-8v-4h11.002v4Z"/>`),
		g.Group(children),
	)
}