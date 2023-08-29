package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDetailsReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28h7v2h-7zm0-4h14v2H16zm0-4h14v2H16zM4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zM28 8H16l-3.414-3.414A2 2 0 0 0 11.172 4H4a2 2 0 0 0-2 2v12h2V6h7.172l3.414 3.414l.586.586H28v8h2v-8a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}