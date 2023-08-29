package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.45.95a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-1.5h1.5a.5.5 0 0 0 0-1h-2Zm4.5 0a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm-.5 12.5a.5.5 0 0 1 .5-.5h3a.5.5 0 1 1 0 1h-3a.5.5 0 0 1-.5-.5Zm-3.5-7.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 1 0v-3Zm11.5-.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Zm-2-4.5a.5.5 0 1 0 0 1h1.5v1.5a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.501-.5H11.45Zm-10 10a.5.5 0 0 1 .5.5v1.5h1.5a.5.5 0 1 1 0 1h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5Zm12.5.5a.5.5 0 0 0-1 0v1.5h-1.5a.5.5 0 1 0 0 1h2a.5.5 0 0 0 .5-.5v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}