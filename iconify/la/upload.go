package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3.594l-.719.687l-7 7L9.72 12.72L15 7.438V24h2V7.437l5.281 5.282l1.438-1.438l-7-7zM7 26v2h18v-2z"/>`),
		g.Group(children),
	)
}