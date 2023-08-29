package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartitionCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 12c-1.858 0-3.41 1.28-3.858 3h-2.728c-.534 0-1.036.208-1.414.586L12.586 21H2v2h10.586c.534 0 1.036-.208 1.414-.586L19.414 17h2.728c.447 1.72 2 3 3.858 3c2.206 0 4-1.794 4-4s-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2zm-9.828-3.414L12.586 11H2V9h10.586c.534 0 1.036.208 1.414.586l3.586 3.586l-1.414 1.414z"/>`),
		g.Group(children),
	)
}