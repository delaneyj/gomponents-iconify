package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8ZM4 4H2v2h1V5h1V4Zm8 1h-1V4h2v2h-1V5ZM7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Zm3.5 5v1h2V9h-1v1h-1ZM2 9h1v1h1v1H2V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}