package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5H0l4-5h12l4 5zm0 2v5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V7h6.126a4.002 4.002 0 0 0 7.748 0H20zM8.265 7h3.465a2 2 0 0 1-3.465 0z"/>`),
		g.Group(children),
	)
}