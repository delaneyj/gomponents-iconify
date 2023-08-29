package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchInteraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 21v-1a1 1 0 0 1 2 0v10h2V20a3.003 3.003 0 0 0-3-3a2.964 2.964 0 0 0-1.47.401a2.954 2.954 0 0 0-4-1A2.993 2.993 0 0 0 19 15a2.96 2.96 0 0 0-1 .185V10a3 3 0 0 0-6 0v11.105l-2.235-1.53v.001a2.999 2.999 0 0 0-3.882 4.55L12.323 30l1.347-1.478l-6.378-5.818A.99.99 0 0 1 7 22a1 1 0 0 1 1.6-.8l5.4 3.695V10a1 1 0 0 1 2 0v11h2v-3a1 1 0 0 1 2 0v3h2v-2a1 1 0 0 1 2 0v2Z"/><path fill="currentColor" d="M28 12h-6v-2h6V4H4v6h4v2H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}