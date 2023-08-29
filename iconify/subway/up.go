package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Up(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 512h139.6V372.4H0V512zm0-186.2h139.6V186.2H0v139.6zm46.5-93.1H93v46.5H46.5v-46.5zM0 139.6h139.6V0H0v139.6zm46.5-93.1H93V93H46.5V46.5zm314.2 35L442.2 0h-256v256l81.5-81.5C453.8 290.9 372.4 418.9 209.5 512c162.9 0 395.6-186.2 151.2-430.5z"/>`),
		g.Group(children),
	)
}