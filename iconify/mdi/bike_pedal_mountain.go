package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BikePedalMountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.82 5a2 2 0 0 0-1.96 1.61l-.82 4.1c-.04.13-.04.26-.04.39v1.8c0 .13 0 .26.04.39l.82 4.1A2 2 0 0 0 6.82 19h6.56a2 2 0 0 0 1.79-1.11l1.62-3.23c.1-.21.17-.43.21-.66h1v-1h2v-2h-2v-1h-1a2.54 2.54 0 0 0-.21-.66l-1.62-3.23A2 2 0 0 0 13.38 5H6.82m0 2h.56l2.5 5l-2.5 5h-.56L6 12.9v-1.8L6.82 7m4.8 4l-2-4h3.76L15 10.24V11h-3.38m0 2H15v.76L13.38 17H9.62l2-4Z"/>`),
		g.Group(children),
	)
}