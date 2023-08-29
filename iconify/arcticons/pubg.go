package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pubg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 12.5h-33c-1.1 0-2 .9-2 2v19c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-19c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.431 28v-8h2.62c1.48 0 2.68 1.203 2.68 2.687s-1.2 2.686-2.68 2.686h-2.62M17.8 20v5.35a2.65 2.65 0 1 0 5.3 0V20m5.318 4a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.3m12.451-1.35a2.65 2.65 0 0 0-2.65-2.65h0a2.65 2.65 0 0 0-2.65 2.65v2.7a2.65 2.65 0 0 0 2.65 2.65h0a2.65 2.65 0 0 0 2.65-2.65h-2.65M42.5 19h1m-1 10h1m-39-10h1m-1 10h1"/>`),
		g.Group(children),
	)
}