package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 21.09v-7.465a6.21 6.21 0 0 0-6.2-6.2H10.7a6.25 6.25 0 0 0-6.2 6.2v7.719"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 18.844a6.21 6.21 0 0 0-6.2-6.2H10.7a6.21 6.21 0 0 0-6.2 6.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.953 21.488a6.212 6.212 0 0 0-5.653-3.656H10.7a6.212 6.212 0 0 0-5.645 3.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.298 27.1L4.5 21.343v13.032a6.21 6.21 0 0 0 6.2 6.2h26.601a6.21 6.21 0 0 0 6.2-6.2V21.09l-5.915 4.302a10.199 10.199 0 0 1-8.288 1.708Z"/>`),
		g.Group(children),
	)
}