package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AriaTwoAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.9 41.9L6.1 6.1m34.4-.6h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.2 34.8c0 1.3-1.1 2.4-2.4 2.4h0c-1.3 0-2.4-1.1-2.4-2.4v-1.6c0-1.3 1.1-2.4 2.4-2.4h0c1.3 0 2.4 1.1 2.4 2.4m0 4v-6.3m1.9-4c0-.8.6-1.4 1.3-1.4s1.4.6 1.4 1.4c0 .4-.2.7-.4 1c-.6.5-2.3 1.8-2.3 1.8h2.7m9.8-9.1c1.1 0 2 .9 2 2.1c0 1.1-.9 2-2.1 2c-1.1 0-2-.9-2-2c.1-1.2 1-2.1 2.1-2.1h0ZM42.5 33H33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.6 14.6c2.8-1.7 6-2.6 9.5-2.6h0c10.1 0 18.4 8.2 18.4 18.4h0m-4.6-18.9l-2.5 4.2"/>`),
		g.Group(children),
	)
}