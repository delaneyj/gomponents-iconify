package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 15V7m0 .5v3m0-3a4 4 0 0 0-4-4h-3v3a4 4 0 0 0 4 4h3m0-3h3a4 4 0 0 0 4-4v-3h-3a4 4 0 0 0-4 4v3Zm0 0l4-4m-4 7l-4-4"/>`),
		g.Group(children),
	)
}