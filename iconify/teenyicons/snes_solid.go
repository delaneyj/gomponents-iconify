package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnesSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.5A1.5 1.5 0 0 1 1.5 3h12A1.5 1.5 0 0 1 15 4.5v6a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 10.5v-6ZM12 7h-1V6h1v1Zm-9 3V9H2V8h1V7h1v1h1v1H4v1H3Zm8-1h1V8h-1v1ZM9 9h1V8H9v1Zm1-2H9V6h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}