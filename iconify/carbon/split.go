package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Split(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 20h2v4h-2zm0-6h2v4h-2zm0-6h2v4h-2z"/><path fill="currentColor" d="M28 16A12.01 12.01 0 0 0 17 4.05V2h-2v2.05a11.99 11.99 0 0 0 0 23.9V30h2v-2.05A12.01 12.01 0 0 0 28 16ZM16 26a10 10 0 1 1 10-10a10.011 10.011 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}