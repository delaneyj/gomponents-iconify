package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0ZM7 5V3.99h1V5H7Zm1 2v3h1v1H6v-1h1V8H6V7h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}