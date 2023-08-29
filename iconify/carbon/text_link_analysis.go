package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextLinkAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 10h8v2H8zm0-4h12v2H8zm0-4h12v2H8z"/><path fill="currentColor" d="m4.711 28l5.631-9.996l7.434 6.49a2 2 0 0 0 3.084-.534l6.97-10.403l-1.661-1.114l-7 10.448l-.07.103l-7.435-6.49a2.003 2.003 0 0 0-3.08.53L4 25.183V2H2v26a2.002 2.002 0 0 0 2 2h26v-2Z"/>`),
		g.Group(children),
	)
}