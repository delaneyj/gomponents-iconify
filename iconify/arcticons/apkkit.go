package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9A2.18 2.18 0 0 0 4.5 11h0v7.29h39v-3.42A1.83 1.83 0 0 0 41.68 13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.28h-39V37a2.18 2.18 0 0 0 2.17 2.2h34.65A2.18 2.18 0 0 0 43.5 37h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.46 32.18A1.23 1.23 0 1 1 28.68 31a1.22 1.22 0 0 1-1.22 1.18Zm7.12 0A1.23 1.23 0 1 1 35.8 31a1.22 1.22 0 0 1-1.22 1.18Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31 24.65h0a8.72 8.72 0 0 1 8.72 8.72v0a1.24 1.24 0 0 1-1.24 1.24h-15a1.24 1.24 0 0 1-1.24-1.24v0A8.72 8.72 0 0 1 31 24.65Zm-7.52-1.3l2.48 2.91m12.58-2.91l-2.47 2.91"/>`),
		g.Group(children),
	)
}