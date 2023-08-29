package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameControllerRetroSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2h12ZM10 6H8V5h2v1ZM4 7V6h1v1h1v1H5v1H4V8H3V7h1Zm8 3h-2V9h2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}