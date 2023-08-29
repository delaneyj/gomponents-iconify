package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.331 14.644L6.537.813l17.55 10.075zM2.938 0c-.813.425-1.356 1.2-1.356 2.206v27.581c0 1.006.544 1.781 1.356 2.206l16.038-16zm26.574 14.1l-3.681-2.131L21.725 16l4.106 4.031l3.756-2.131c1.125-.893 1.125-2.906-.075-3.8zM6.538 31.188l17.55-10.075l-3.756-3.756z"/>`),
		g.Group(children),
	)
}