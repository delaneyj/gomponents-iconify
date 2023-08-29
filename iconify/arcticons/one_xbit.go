package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneXbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m10.76 2.091l4.618 4.618m0-4.618l-4.618 4.618"/><circle cx="33.343" cy="19.058" r=".851" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.343 22.035v6.441m3.577-8.446v7.231a1.215 1.215 0 0 0 1.215 1.215h.365m-2.856-6.441h2.552M28.83 24a2.243 2.243 0 0 1 0 4.486h-3.7v-8.972h3.7a2.243 2.243 0 0 1 0 4.486Zm.001 0H25.13m0-4.486h-1.315m1.315 8.972h-1.315m2.456-8.972v-1.59m1.831 1.59v-1.59m-1.831 12.152v-1.59m1.831 1.59v-1.59M9.5 28.861h4.861M9.5 20.462l2.431-1.323m0 0v9.722"/>`),
		g.Group(children),
	)
}