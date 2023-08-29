package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nextcloudbookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm0 11.18l3.59 6.54l7.18 1.51l-5 5.51l.82 7.46L24 31.56l-6.68 3.1l.87-7.45l-5-5.54l7.2-1.46L24 13.68Z"/>`),
		g.Group(children),
	)
}