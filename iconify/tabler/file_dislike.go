package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDislike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 15a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm3 0a1 1 0 0 1 1-1h3.756a1 1 0 0 1 .958.713l1.2 3c.09.303.133.63-.056.884c-.188.254-.542.403-.858.403h-2v2.467a1.1 1.1 0 0 1-2.015.61L6 19v-4zm8-12v4a1 1 0 0 0 1 1h4"/><path d="M5 11V5a2 2 0 0 1 2-2h7l5 5v11a2 2 0 0 1-2 2h-2.5"/></g>`),
		g.Group(children),
	)
}