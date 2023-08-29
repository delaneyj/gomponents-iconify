package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartitionSame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18c-1.858 0-3.41 1.28-3.858 3H2v2h20.142c.447 1.72 2 3 3.858 3c2.206 0 4-1.794 4-4s-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2zm0-18c-1.858 0-3.41 1.28-3.858 3H2v2h20.142c.447 1.72 2 3 3.858 3c2.206 0 4-1.794 4-4s-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2z"/>`),
		g.Group(children),
	)
}