package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eruditedictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.807 31.025A12.903 12.903 0 1 1 36.887 24H2.522A21.483 21.483 0 1 1 5.4 34.748"/>`),
		g.Group(children),
	)
}