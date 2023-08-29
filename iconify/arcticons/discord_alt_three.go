package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscordAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.59 34.173A48.874 48.874 0 0 1 14.68 38C7.3 37.79 4.5 33 4.5 33a44.83 44.83 0 0 1 4.81-19.52A16.47 16.47 0 0 1 18.69 10l1 2.31A32.693 32.693 0 0 1 24 12a33.01 33.01 0 0 1 4.33.3l1-2.31a16.47 16.47 0 0 1 9.38 3.51A44.83 44.83 0 0 1 43.5 33s-2.8 4.79-10.18 5a47.42 47.42 0 0 1-2.86-3.81m6.46-2.9c-3.84 1.945-7.556 3.89-12.92 3.89s-9.08-1.945-12.92-3.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.197 26.23a3.35 3.35 0 1 1-6.7 0m19.006 0a3.35 3.35 0 1 1-6.7 0"/>`),
		g.Group(children),
	)
}