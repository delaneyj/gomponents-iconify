package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NonCertified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 8h10v2H6zm0 4h8v2H6zm0 4h4v2H6z"/><path fill="currentColor" d="M28 26H7.414L30 3.414L28.586 2l-2 2H4a2.002 2.002 0 0 0-2 2v16h2V6h20.586L2 28.586L3.414 30l2-2H28a2.002 2.002 0 0 0 2-2V10h-2Z"/>`),
		g.Group(children),
	)
}