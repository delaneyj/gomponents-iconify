package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.02 4.975A9.93 9.93 0 0 0 2.07 12A9.935 9.935 0 0 0 12 21.935a9.98 9.98 0 0 0 3.8-.75a10.189 10.189 0 0 0 3.22-2.16a9.934 9.934 0 0 0 0-14.05Zm-.7 13.34a8.921 8.921 0 0 1-5.32 2.57v-6.56h1.88a1 1 0 0 0 0-2H13v-2.74a1 1 0 0 1 1-1h1.2a1 1 0 0 0 0-2h-1.7a2.5 2.5 0 0 0-2.5 2.5v3.24H9.13a1 1 0 1 0 0 2H11v6.56a8.919 8.919 0 1 1 9.26-5.47a9.061 9.061 0 0 1-1.94 2.9Z"/>`),
		g.Group(children),
	)
}