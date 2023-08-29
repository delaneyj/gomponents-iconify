package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowModelerReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30h-8v-8h8zm-6-2h4v-4h-4zM4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm15-9h-6l-3 4l6 6l6-6z"/><path fill="currentColor" d="M24 17v-2a4 4 0 0 0 0-8H12V5h12a6 6 0 0 1 0 12zm-14-7H2V2h8zM4 8h4V4H4z"/>`),
		g.Group(children),
	)
}