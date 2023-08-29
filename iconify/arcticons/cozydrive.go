package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cozydrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.25 34.43h-8.8a3.71 3.71 0 1 1 0-7.42h0a3.35 3.35 0 0 1 .71.07v-.32a3.71 3.71 0 0 1 7.42 0v.32a4.17 4.17 0 0 1 .71-.07a3.71 3.71 0 1 1 0 7.42Z"/><path fill="currentColor" d="M34.63 29.9a5.39 5.39 0 0 1-1.78.33a4.87 4.87 0 0 1-1.77-.33a.44.44 0 0 0-.59.23a.43.43 0 0 0 .22.58a5.76 5.76 0 0 0 2.13.41a5.83 5.83 0 0 0 2.16-.4a.45.45 0 0 0 .17-.61a.44.44 0 0 0-.53-.21Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9A2.18 2.18 0 0 0 4.5 11h0v7.29h39v-3.42A1.83 1.83 0 0 0 41.68 13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.28h-39V37a2.18 2.18 0 0 0 2.17 2.2h34.65A2.18 2.18 0 0 0 43.5 37h0Z"/>`),
		g.Group(children),
	)
}