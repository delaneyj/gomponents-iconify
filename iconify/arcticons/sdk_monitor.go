package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdkMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.56 30.4V17.6h2.88a5.6 5.6 0 0 1 5.6 5.6v1.6a5.6 5.6 0 0 1-5.6 5.6Zm11.06-12.8v12.8m6.88 0L33.23 24l5.27-6.357M33.23 24h-1.61M9.754 28.998a3.58 3.58 0 0 0 3.138 1.402h1.895a3.197 3.197 0 0 0 3.193-3.2h0a3.197 3.197 0 0 0-3.193-3.2h-2.094A3.197 3.197 0 0 1 9.5 20.8h0a3.197 3.197 0 0 1 3.193-3.2h1.895a3.58 3.58 0 0 1 3.139 1.402"/>`),
		g.Group(children),
	)
}