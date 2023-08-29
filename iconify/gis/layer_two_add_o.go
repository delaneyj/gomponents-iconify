package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerTwoAddO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M80.5 10.797c-7.145 0-13.252 4.246-15.977 10.357H28.135a3.5 3.5 0 0 0-2.668 1.235L.832 51.404c-1.93 2.275-.313 5.767 2.67 5.766l93-.065c2.981-.002 4.595-3.492 2.666-5.765l-7.87-9.272A17.43 17.43 0 0 0 98 28.297c0-9.695-7.805-17.5-17.5-17.5zm-3.5 5h7v9h9v7h-9v9h-7v-9h-9v-7h9zM29.754 28.154h33.254c0 .048-.008.095-.008.143c0 9.694 7.805 17.5 17.5 17.5c1.492 0 2.93-.205 4.31-.553l4.131 4.867l-77.873.053zM89.91 62.11l-9.178.008l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 73.08c-1.93 2.274-.313 5.767 2.67 5.766l93-.065c2.981-.002 4.595-3.492 2.666-5.765z" color="currentColor"/>`),
		g.Group(children),
	)
}