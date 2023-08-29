package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Operation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M23 26h-2l5-10l-5-10h2l5 10l-5 10z" fill="currentColor"/><path d="M4 6h2v20H4z" fill="currentColor"/><path d="M16 9h-4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V11a2 2 0 0 0-2-2zm0 12h-4V11h4z" fill="currentColor"/><path d="M13 15h2v2h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}