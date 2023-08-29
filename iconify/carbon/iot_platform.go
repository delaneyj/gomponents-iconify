package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IotPlatform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 19h-4v-4h-2v9H8V8h9V6h-4V2h-2v4H8a2.002 2.002 0 0 0-2 2v3H2v2h4v6H2v2h4v3a2.002 2.002 0 0 0 2 2h3v4h2v-4h6v4h2v-4h3a2.003 2.003 0 0 0 2-2v-3h4Z"/><path fill="currentColor" d="M21 21H11V11h10zm-8-2h6v-6h-6zm18-6h-2A10.012 10.012 0 0 0 19 3V1a12.013 12.013 0 0 1 12 12z"/><path fill="currentColor" d="M26 13h-2a5.006 5.006 0 0 0-5-5V6a7.008 7.008 0 0 1 7 7Z"/>`),
		g.Group(children),
	)
}