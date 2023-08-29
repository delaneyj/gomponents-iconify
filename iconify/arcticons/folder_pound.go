package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M41.68 13H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9A2.18 2.18 0 0 0 4.5 11h0v7.29h39v-3.42A1.83 1.83 0 0 0 41.71 13h-.03Z"/><path d="M4.5 18.28V37a2.18 2.18 0 0 0 2.16 2.2h34.66a2.18 2.18 0 0 0 2.18-2.18V37h0V18.28"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.225 25.904a3.15 3.15 0 0 0-3.15-3.15h0a3.15 3.15 0 0 0-3.15 3.15v5.666a3 3 0 0 1-.61 1.813l-1.04 1.37h7.95m-4.65-5.587h-3.3"/>`),
		g.Group(children),
	)
}