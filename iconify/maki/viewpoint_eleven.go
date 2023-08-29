package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewpointEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.403 5.777a.852.852 0 0 0-.183.215L.75 3.322a5.78 5.78 0 0 1 1.1-1.11zM9.15 2.222l-2.544 3.56a1.379 1.379 0 0 1 .184.2l3.47-2.66a5.78 5.78 0 0 0-1.11-1.1zM5.37 7.009a1 1 0 1 0 1.122.86a1 1 0 0 0-1.122-.86zm-.65-5.987a5.774 5.774 0 0 0-1.52.41l1.768 4.003a.815.815 0 0 1 .238-.061zm1.56.01l-.454 4.353a.761.761 0 0 1 .206.044L7.79 1.431a5.519 5.519 0 0 0-1.51-.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}