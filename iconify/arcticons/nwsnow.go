package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nwsnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2ZM5.5 24h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 19.03v-8.56l5.67 8.56v-8.56m10.61 0l-2.14 8.56L24 10.47l-2.14 8.56l-2.14-8.56m10.78 7.62a2.37 2.37 0 0 0 2.1.94h1.26A2.14 2.14 0 0 0 36 16.89h0a2.14 2.14 0 0 0-2.14-2.14h-1.4a2.13 2.13 0 0 1-2.13-2.14h0a2.13 2.13 0 0 1 2.13-2.14h1.27a2.37 2.37 0 0 1 2.1.94M12 37.3v-8.56l5.67 8.56v-8.56M36 29.2l-2.14 8.56l-2.14-8.56l-2.14 8.56l-2.14-8.56m-7.57 5.27a2.84 2.84 0 0 0 5.67 0v-2.89a2.84 2.84 0 1 0-5.67 0Z"/>`),
		g.Group(children),
	)
}