package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscordAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.54 34.22A47.42 47.42 0 0 1 14.68 38C7.3 37.79 4.5 33 4.5 33a44.83 44.83 0 0 1 4.81-19.52A16.47 16.47 0 0 1 18.69 10l1 2.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.85 22.67a3.48 3.48 0 0 0-3.37 3.9a3.38 3.38 0 0 0 3.31 3.22a3.53 3.53 0 0 0 3.43-3.9a3.45 3.45 0 0 0-3.37-3.22Zm-5.65-8.3a28.19 28.19 0 0 1 8.16-2.18A23.26 23.26 0 0 1 24 12a23.26 23.26 0 0 1 3.64.21a28.19 28.19 0 0 1 8.16 2.18m-7.47-2.09l1-2.31a16.47 16.47 0 0 1 9.38 3.51A44.83 44.83 0 0 1 43.5 33s-2.8 4.79-10.18 5a47.42 47.42 0 0 1-2.86-3.81m6.46-2.9a29.63 29.63 0 0 1-8.64 3.49a21.25 21.25 0 0 1-4.28.4h0a21.25 21.25 0 0 1-4.28-.4a29.63 29.63 0 0 1-8.64-3.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.15 22.67a3.48 3.48 0 0 1 3.37 3.9a3.38 3.38 0 0 1-3.31 3.22a3.53 3.53 0 0 1-3.43-3.9a3.45 3.45 0 0 1 3.37-3.22Z"/>`),
		g.Group(children),
	)
}