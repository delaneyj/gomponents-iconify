package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.291 27.547h3.667m-3.667-7.334h3.667m-3.667 3.667h2.383m-2.383-3.667v7.334m8.112-.143h3.667m-3.667-7.335h3.667m-3.667 3.667h2.383m-2.383-3.667v7.335m-16.575-7.388h4.859l-4.859 7.334h4.859m19.146-.485a2.876 2.876 0 0 0 2.2.642h.275a2.427 2.427 0 0 0 2.384-2.384h0a2.427 2.427 0 0 0-2.384-2.384h-2.475v-2.567h4.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.951 36.955A13.189 13.189 0 0 1 25.642 11.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.637 40.992a17.935 17.935 0 1 1 1.55-32.417"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.405 45.5a21.502 21.502 0 0 1 .277-43"/>`),
		g.Group(children),
	)
}