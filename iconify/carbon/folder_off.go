package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h-2.586L30 3.414L28.586 2L2 28.586L3.414 30l2-2H28a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2zm0 18H7.414l16-16H28zM4 6h7.172l3.414 3.414l.586.586H18V8h-2l-3.414-3.414A2 2 0 0 0 11.172 4H4a2 2 0 0 0-2 2v18h2z"/>`),
		g.Group(children),
	)
}