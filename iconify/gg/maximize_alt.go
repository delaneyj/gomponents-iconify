package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h6v2H6.462l4.843 4.843l-1.415 1.414L5 6.367V9H3V3Zm0 18h6v-2H6.376l4.929-4.928l-1.415-1.414L5 17.548V15H3v6Zm12 0h6v-6h-2v2.524l-4.867-4.866l-1.414 1.414L17.647 19H15v2Zm6-18h-6v2h2.562l-4.843 4.843l1.414 1.414L19 6.39V9h2V3Z"/>`),
		g.Group(children),
	)
}