package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MijnDsw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.508 5.508h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.5 18.905l-2.548 10.19l-2.547-10.19l-2.548 10.19l-2.547-10.19m-8.937 9.075c.625.812 1.409 1.115 2.5 1.115h1.508a2.542 2.542 0 0 0 2.542-2.543h0v-.01A2.542 2.542 0 0 0 23.381 24h-1.664a2.545 2.545 0 0 1-2.545-2.545h0a2.55 2.55 0 0 1 2.55-2.55h1.5c1.09 0 1.874.303 2.5 1.117M9.5 29.095v-10.19h2.293a4.458 4.458 0 0 1 4.458 4.458v1.274a4.458 4.458 0 0 1-4.458 4.458H9.5Z"/>`),
		g.Group(children),
	)
}