package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavycheckmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#2B3B47" d="m484.065 66.324l-71.534-41.759c-8.811-5.143-20.123-2.17-25.266 6.641L211.883 331.642L96.414 264.236c-7.143-4.17-16.314-1.76-20.484 5.384l-45.284 77.573c-4.17 7.143-1.76 16.314 5.384 20.484l205.968 120.235c7.143 4.17 16.314 1.76 20.484-5.384l1.669-2.86c.034-.056.074-.106.107-.163L490.706 91.59c5.143-8.811 2.17-20.123-6.641-25.266z"/>`),
		g.Group(children),
	)
}