package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobilizon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.82 7.67c-.08 2.87-.63 3.55-3.66 3.54c-2.88 0-4-.62-3.94-3.48c0-2.65.78-3.1 3.72-3.15c3.12-.04 3.95.67 3.88 3.09ZM36 29.16c.05 8.66-4.83 14.54-12.18 14.42c-6.53-.09-11.82-5-11.82-14.42s4.46-14 11.94-14.29S36 20.15 36 29.16Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.28 29.29c0 6-1.49 9.31-5.4 9.29c-3.49 0-5.11-3.4-5.13-9.28c0-6.18 2-9.19 5.23-9.29c3-.1 5.35 2.74 5.31 9.28Z"/>`),
		g.Group(children),
	)
}