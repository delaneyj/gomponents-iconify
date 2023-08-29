package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZLpar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 10h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1h-3a2 2 0 0 0-2 2v8h-3V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V17h3v8a2 2 0 0 0 2 2h3v1a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1h-3v-8h3v1a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1h-3V7h3v1a2 2 0 0 0 2 2Zm0-6h4v4h-4V4Zm-12.6 8L4 17.92V6.08L11.4 12Zm.6 2.08l.001 11.84l-7.4-5.92l7.4-5.92Zm0-4.16L5.85 5H12v4.92ZM4 22.08L10.15 27H4v-4.92ZM24 24h4v4h-4v-4Zm0-10h4v4h-4v-4Z"/>`),
		g.Group(children),
	)
}