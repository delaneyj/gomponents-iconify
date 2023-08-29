package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.38 5.59a3.69 3.69 0 1 0-3.69 3.69a3.67 3.67 0 0 0 2.483-.976L9 9.991l.012.009l-.004.003l-1.836 1.693a3.665 3.665 0 0 0-2.482-.976a3.69 3.69 0 1 0 3.69 3.69c0-.297-.044-.582-.111-.858l2.844-1.991l4.127 3.065c2.212 1.549 3.76-.663 3.76-.663L8.269 6.448c.066-.276.111-.561.111-.858zm-3.69 1.8a1.8 1.8 0 1 1 0-3.6a1.8 1.8 0 0 1 0 3.6zm0 8.82a1.8 1.8 0 1 1 0-3.6a1.8 1.8 0 0 1 0 3.6zM19 6.038s-1.548-2.212-3.76-.663L12.035 7.61l2.354 1.648L19 6.038z"/>`),
		g.Group(children),
	)
}