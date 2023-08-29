package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Communication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.5 6C10.492 6 4 12.492 4 20.5C4 38.5 28 42 28 42v-7h1.5C37.508 35 44 28.508 44 20.5S37.508 6 29.5 6h-11ZM24 23.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM34.5 21a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM16 23.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}