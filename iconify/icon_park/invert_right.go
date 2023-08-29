package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 15C38 22.2989 33.897 28.5832 28 31.4081C25.8653 32.4307 23.4954 33 21 33C11.6112 33 4 24.9411 4 15"/><path d="M30 20L38 15L44 23"/></g>`),
		g.Group(children),
	)
}