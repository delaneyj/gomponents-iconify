package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroidFortyEightSx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.4 19.7h24.2a1.69 1.69 0 0 1 1.8 1.6v12.1a1.69 1.69 0 0 1-1.8 1.6H12.4a1.69 1.69 0 0 1-1.8-1.6v-12a1.77 1.77 0 0 1 1.8-1.7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.4 24.4H27a11.58 11.58 0 0 0-3.8 7.9m-10.6-8.2c0-1.2.3-2.3 1.4-2.4m22.4 2.4c0-1.2-.3-2.3-1.4-2.4M24.4 9v7.5h3.8m6.9 0h3.8M35.1 9h3.8m-3.8 3.8h2.5M35.1 9v7.5m-2-7.5l-2.5 7.5L28.2 9m-17.9 6.7a2.06 2.06 0 0 0 1.8.8h1.1a1.9 1.9 0 0 0 1.9-1.9h0a1.9 1.9 0 0 0-1.9-1.9H12a1.9 1.9 0 0 1-1.9-1.9h0A2 2 0 0 1 12 9h1.1a2 2 0 0 1 1.8.8m2.3 4.2a2.5 2.5 0 1 0 5 0v-2.5a2.5 2.5 0 0 0-5 0Z"/>`),
		g.Group(children),
	)
}