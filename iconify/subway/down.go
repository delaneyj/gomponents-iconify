package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Down(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 512h139.6V372.4H0V512zm46.5-93.1H93v46.5H46.5v-46.5zM0 139.6h139.6V0H0v139.6zm0 186.2h139.6V186.2H0v139.6zm46.5-93.1H93v46.5H46.5v-46.5zM209.5 0c162.9 93.1 244.4 221.1 58.2 337.5L186.2 256v256h256l-81.5-81.5C605.1 186.2 372.4 0 209.5 0z"/>`),
		g.Group(children),
	)
}