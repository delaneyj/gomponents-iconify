package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 176v-64h-64V69h64V5h43v64h64v43h-64v64h-43zm-85 192q18 0 30.5 12.5t12.5 30t-12.5 30T128 453t-30-12.5t-12-30t12-30t30-12.5zm213.5 0q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5zM132 299q0 5 5 5h247v43H128q-18 0-30.5-12.5T85 304q0-11 6-20l28-53L43 69H0V27h70l20 42l20 43l48 101l3 6h149l59-107l24-43l37 21l-82 149q-12 22-38 22H151l-19 35v3z"/>`),
		g.Group(children),
	)
}