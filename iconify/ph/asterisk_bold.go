package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M218.29 182.17a12 12 0 0 1-16.47 4.12L140 149.19V216a12 12 0 0 1-24 0v-66.81l-61.82 37.1a12 12 0 1 1-12.35-20.58L104.68 128L41.83 90.29a12 12 0 1 1 12.35-20.58l61.82 37.1V40a12 12 0 0 1 24 0v66.81l61.82-37.1a12 12 0 1 1 12.35 20.58L151.32 128l62.85 37.71a12 12 0 0 1 4.12 16.46Z"/>`),
		g.Group(children),
	)
}