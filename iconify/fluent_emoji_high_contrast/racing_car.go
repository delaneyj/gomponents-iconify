package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RacingCar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 19v4.671a4.491 4.491 0 0 0-4-1.643V19h4Zm-7.242 10a4.478 4.478 0 0 1-.73-3h-9.055a4.478 4.478 0 0 1-.73 3h10.515ZM8.5 30a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-2a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm18 2a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-2a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM21 24h1.758a4.483 4.483 0 0 0-.502 1h-9.512a4.505 4.505 0 0 0-1.207-1.82A16.96 16.96 0 0 1 14 23h2l1 1h2v-2c0-.55.45-1 1-1a7.24 7.24 0 0 1 3.24.764L21 24ZM4 26h.027a4.479 4.479 0 0 0 .73 3H2v-2l2-1Z"/>`),
		g.Group(children),
	)
}