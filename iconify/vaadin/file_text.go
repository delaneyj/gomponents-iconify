package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0v4h4z"/><path fill="currentColor" d="M9 0H2v16h12V5H9V0zm3 12H4v-1h8v1zm0-2H4V9h8v1zm0-3v1H4V7h8z"/>`),
		g.Group(children),
	)
}