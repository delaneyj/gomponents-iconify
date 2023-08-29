package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IotConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 19h-4v-4h-2v9H8V8h9V6h-4V2h-2v4H8a2.002 2.002 0 0 0-2 2v3H2v2h4v6H2v2h4v3a2.002 2.002 0 0 0 2 2h3v4h2v-4h6v4h2v-4h3a2.003 2.003 0 0 0 2-2v-3h4Z"/><path fill="currentColor" d="M26 2a4.004 4.004 0 0 0-4 4a3.956 3.956 0 0 0 .567 2.02L19.586 11H11v10h10v-8.586l2.98-2.98A3.955 3.955 0 0 0 26 10a4 4 0 0 0 0-8Zm-7 17h-6v-6h6Zm7-11a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}