package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feArtboard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArtboard1" fill="currentColor"><path id="feArtboard2" d="M6 16V8H4a1 1 0 1 1 0-2h2V4a1 1 0 1 1 2 0v2h8V4a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2h-2v8h2a1 1 0 0 1 0 2h-2v2a1 1 0 0 1-2 0v-2H8v2a1 1 0 0 1-2 0v-2H4a1 1 0 0 1 0-2h2Zm2 0h8V8H8v8Z"/></g></g>`),
		g.Group(children),
	)
}