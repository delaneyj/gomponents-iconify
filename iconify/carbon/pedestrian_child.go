package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PedestrianChild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 16h-8.532l-5-6H5a3.003 3.003 0 0 0-3 3v6a2.002 2.002 0 0 0 2 2v7a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2v-7h-2v7H6v-9H4v-6a1 1 0 0 1 1-1h6.532l5 6H26a1 1 0 0 1 1 1v3h-2v6h-3v-6h-2v6a2.002 2.002 0 0 0 2 2h3a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0 2-2v-3a3.003 3.003 0 0 0-3-3Z"/><path fill="currentColor" d="M23.5 15a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5zM8 9a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}