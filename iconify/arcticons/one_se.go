package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneSe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 31.03h7.082M9.5 18.814l3.541-1.948m0 0V31.03m6.223-1.5a4.25 4.25 0 0 0 3.565 1.604h2.14a3.576 3.576 0 0 0 3.565-3.565h0a3.576 3.576 0 0 0-3.566-3.565h-2.317a3.576 3.576 0 0 1-3.565-3.566h0a3.576 3.576 0 0 1 3.565-3.565h2.139a3.828 3.828 0 0 1 3.565 1.604m3.063 12.553H38.5m-7.082-14.164H38.5m-7.082 7.082h4.603m-4.603-7.082V31.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}