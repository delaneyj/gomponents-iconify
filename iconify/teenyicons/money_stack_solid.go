package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyStackSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11h12A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-12ZM4 2H2v2h1V3h1V2Zm3.5 1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM12 3h-1V2h2v2h-1V3ZM3 7H2v2h2V8H3V7Zm8 2V8h1V7h1v2h-2Z" clip-rule="evenodd"/><path fill="currentColor" d="M0 12v1h15v-1H0Zm0 2v1h15v-1H0Z"/>`),
		g.Group(children),
	)
}