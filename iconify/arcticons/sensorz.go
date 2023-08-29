package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sensorz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.1 26.8h0a1.47 1.47 0 0 1-1.7-1.5v-1.1a2 2 0 0 1 1.7-1.9h0a1.47 1.47 0 0 1 1.7 1.5v1.1a2 2 0 0 1-1.7 1.9Zm3.3-3.3a2 2 0 0 1 1.7-1.9h0m-1.7.2v4.5m2.6-7.1l4.5-.6L35 26l4.5-.6M8.6 29.2a1.79 1.79 0 0 0 1.4.2l.4-.1a1.38 1.38 0 0 0 1.1-1.3h0a1 1 0 0 0-1.1-1l-.8.1a1 1 0 0 1-1.1-1h0a1.38 1.38 0 0 1 1.1-1.3l.4-.1a1.89 1.89 0 0 1 1.4.2M23 27.2a1.79 1.79 0 0 0 1.4.2l.4-.1a1.38 1.38 0 0 0 1.1-1.3h0a1 1 0 0 0-1.1-1l-.8.2a1 1 0 0 1-1.1-1h0a1.38 1.38 0 0 1 1.1-1.3l.4-.1a1.89 1.89 0 0 1 1.4.2m-9.6 4.7a2.17 2.17 0 0 1-1.5 1.1h0a1.47 1.47 0 0 1-1.7-1.5v-1.1a2 2 0 0 1 1.7-1.9h0a1.47 1.47 0 0 1 1.7 1.5v.5l-3.4.4m8.3 1.1V25a1.53 1.53 0 0 0-1.7-1.5h0a2 2 0 0 0-1.7 1.9v2.8m0-2.8v-1.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}