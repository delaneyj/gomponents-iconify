package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepoPull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 1.5V3h-1V2H3v8h10v3.5l-.5.5H8v-1h4v-2H2.735a.72.72 0 0 0-.285.06a.74.74 0 0 0-.4.4a.93.93 0 0 0-.05.29v.5a.93.93 0 0 0 .05.29a.74.74 0 0 0 .4.4c.091.04.19.06.29.06H3v1h-.26a1.9 1.9 0 0 1-.67-.13a1.77 1.77 0 0 1-.94-.95a1.7 1.7 0 0 1-.13-.67v-9.5a1.7 1.7 0 0 1 .13-.62a1.77 1.77 0 0 1 .94-1A1.9 1.9 0 0 1 2.74 1h9.76l.5.5zM2 10.17V2.748v7.422zM5 3H4v1h1V3zm0 2H4v1h1V5zM4 7h1v1H4V7zm8.07-3.61l-.7.71l1.92 1.92H7v1h6.39l-2.02 2.03l.7.7l2.83-2.82v-.71l-2.83-2.83zM5.5 13.49L4.28 15H4v-3h3v3h-.28L5.5 13.49z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}