package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 24V6H1a3 3 0 0 1 3 3v4V3h.5a3 3 0 0 1 3 3v6.5h2a3 3 0 0 1 3 3v6h-2A2.5 2.5 0 0 0 8 24m1-8.5V19m0-1.75h3.5M23.5 0v18H23a3 3 0 0 1-3-3v-4v10h-.5a3 3 0 0 1-3-3v-6.5h-2a3 3 0 0 1-3-3v-6h2A2.5 2.5 0 0 0 16 0m-1 8.5V5m0 1.75h-3.5"/>`),
		g.Group(children),
	)
}