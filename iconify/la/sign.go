package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v2H4v2h2v20h2V8h20V6H8V4H6zm4 6v12h16V10H10zm2 2h12v8H12v-8z"/>`),
		g.Group(children),
	)
}