package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Injection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M38.168 22.262L19.077 41.354L6.349 28.626L25.44 9.534" clip-rule="evenodd"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M38.168 22.262L19.077 41.354L6.349 28.626L25.44 9.534"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m21.905 5.999l19.8 19.799m-26.871 2.828l4.243 4.243M6.35 41.353l6.363-6.363m19.092-19.092l3.534-3.535"/></g>`),
		g.Group(children),
	)
}