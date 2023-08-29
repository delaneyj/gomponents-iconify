package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloggerB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 7a6 6 0 0 0-6 6v7a6 6 0 0 0 6 6h7a6 6 0 0 0 6-6v-5a1 1 0 0 0-1-1h-1a1 1 0 0 1-1-1a6 6 0 0 0-6-6h-4zm0 5h4c.55 0 1 .45 1 1s-.45 1-1 1h-4c-.55 0-1-.45-1-1s.45-1 1-1zm0 7h7c.55 0 1 .45 1 1s-.45 1-1 1h-7c-.55 0-1-.45-1-1s.45-1 1-1z"/>`),
		g.Group(children),
	)
}