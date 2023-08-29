package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M165 85.3v325.3c-9.7-3.3-20.5-5.3-32-5.3c-41.2 0-74.7 23.9-74.7 53.3c0 29.5 33.4 53.3 74.7 53.3c41.2 0 74.7-23.9 74.7-53.3c0-3.6-.5-7.2-1.5-10.7h1.5V179.8L421 118.9v227.8c-9.7-3.3-20.5-5.3-32-5.3c-41.2 0-74.7 23.9-74.7 53.3c0 29.5 33.4 53.3 74.7 53.3c41.2 0 74.7-23.9 74.7-53.3c0-3.4-.5-6.7-1.4-10l1.4-.7V0L165 85.3z"/>`),
		g.Group(children),
	)
}