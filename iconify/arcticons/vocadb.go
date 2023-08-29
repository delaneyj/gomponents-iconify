package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vocadb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2ZM32 24h10.5m-26.9 0h11.7M5.5 24h2.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.4 17.3h-3.1v13.4h4l.7-1V18.8l-1.6-1.5zM15.6 24l-1.9 6.7l-3.8-13.4l-2 6.7m8.9 1.5h1.5v4.2h-1.5zm20.8-2.9l-3.8-5.3v13.4h4.4l1.9-2.3l-3.2-4.4l.7-1.4zm-16.1 2.9h-1l-.8 2.1l.9 2.1h.9m3.9 0l-1.3-4.2l-1.3 3.7h1.1"/>`),
		g.Group(children),
	)
}