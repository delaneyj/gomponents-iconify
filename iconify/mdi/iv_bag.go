package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IvBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 1h-4C5 1 5 3 5 3v12s0 2 5 2v2h1v1c0 1.21.8 3 3 3h4v-2h-4c-.81 0-1-.55-1-1v-1h1v-2c5 0 5-2 5-2V3s0-2-5-2m3 11h-3v-1h3v1m0-7h-3v1h3v2h-3v1h3v1H7V3.5c.3-.18 1.13-.5 3-.5h4c1.88 0 2.7.32 3 .5V5Z"/>`),
		g.Group(children),
	)
}