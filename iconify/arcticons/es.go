package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Es(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9A2.18 2.18 0 0 0 4.5 11h0v7.29h39v-3.42A1.83 1.83 0 0 0 41.68 13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.28h-39V37a2.18 2.18 0 0 0 2.17 2.2h34.65A2.18 2.18 0 0 0 43.5 37h0ZM25.32 33.82h5.1m-5.1-10.2h5.1m-5.1 5.1h3.33m-3.33-5.1v10.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.58 32.7a2.85 2.85 0 0 0 2.5 1.12h1.51a2.55 2.55 0 0 0 2.54-2.55h0a2.55 2.55 0 0 0-2.54-2.55h-1.67a2.55 2.55 0 0 1-2.55-2.55h0a2.55 2.55 0 0 1 2.55-2.55h1.51a2.86 2.86 0 0 1 2.5 1.11"/>`),
		g.Group(children),
	)
}