package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidSambaClient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.68 13.015H24.77c-2-.1-5.93-4.23-8.19-4.23h-9.9a2.18 2.18 0 0 0-2.18 2.18v.05h0v7.29h39v-3.42a1.83 1.83 0 0 0-1.79-1.87Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 31.94V18.294h-39v18.72a2.18 2.18 0 0 0 2.16 2.2h15.805"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.756 24.694a7.723 7.723 0 0 1 8.001 6.177a5.175 5.175 0 0 1 .038 10.308H26.338a6.226 6.226 0 0 1 0-12.38a7.602 7.602 0 0 1 6.418-4.105Z"/>`),
		g.Group(children),
	)
}