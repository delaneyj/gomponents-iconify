package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkThreeReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm26 10h-8v-8h8zm-6-2h4v-4h-4zm-4-8h-8v-8h8zm-6-2h4v-4h-4z"/><path fill="currentColor" d="M24 17h-2v-2h2a4 4 0 0 0 0-8H12V5h12a6 6 0 0 1 0 12zm-14-7H2V2h8zM4 8h4V4H4z"/>`),
		g.Group(children),
	)
}