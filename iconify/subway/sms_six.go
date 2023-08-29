package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M257 0C115.6 0 1 85.9 1 192c0 75 57.5 139.8 141.1 171.4L86.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S398.4 0 257 0zm128 213.3H278.3V320h-42.7V213.3H129v-42.7h106.7V64h42.7v106.7H385v42.6z"/>`),
		g.Group(children),
	)
}