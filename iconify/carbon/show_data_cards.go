package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowDataCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 10H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 4v4h24V4zm24 26H4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 24v4h24v-4zm24-4H4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 14v4h24v-4z"/>`),
		g.Group(children),
	)
}