package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLegend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.496 5.086a2.084 2.084 0 0 0-.03.002a2.084 2.084 0 0 0-1.77 1.04L.278 24.173a2.084 2.084 0 0 0 1.805 3.125h20.834a2.084 2.084 0 0 0 1.803-3.125L14.305 6.129a2.084 2.084 0 0 0-1.809-1.043zM40 14.486v7h15v-7H40zm22 0v7h38v-7H62zM2.084 39.672A2.084 2.084 0 0 0 0 41.756v16.666a2.084 2.084 0 0 0 2.084 2.084h20.832A2.084 2.084 0 0 0 25 58.422V41.756a2.084 2.084 0 0 0-2.084-2.084H2.084zM40 47.838v7h27v-7H40zm34 0v7h26v-7H74zM12.5 69.914c-6.879 0-12.5 5.621-12.5 12.5s5.621 12.5 12.5 12.5S25 89.293 25 82.414s-5.621-12.5-12.5-12.5zM40 81.19v7h15v-7H40zm22 0v7h38v-7H62z" color="currentColor"/>`),
		g.Group(children),
	)
}