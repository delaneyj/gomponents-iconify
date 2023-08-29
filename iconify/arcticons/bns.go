package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.44 25.95A5.56 5.56 0 0 1 23 20.39h0a5.56 5.56 0 0 1 5.56 5.56v3.61A5.56 5.56 0 0 1 23 35.12h0a5.56 5.56 0 0 1-5.56-5.56m0 5.56V12.88"/><circle cx="31.15" cy="35.12" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}