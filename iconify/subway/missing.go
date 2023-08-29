package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Missing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M499.5 385.4L308.9 57.2c-31.8-52.9-74.1-52.9-105.9 0L12.5 385.4c-31.8 52.9 0 95.3 63.5 95.3h360c63.5 0 95.3-42.4 63.5-95.3zm-201.1 52.9h-84.7v-84.7h84.7v84.7zm0-127h-84.7V120.7h84.7v190.6z"/>`),
		g.Group(children),
	)
}