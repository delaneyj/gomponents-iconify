package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L14.11 16H13l-.28-1.39l-.61-.61H7v7H5V6.89L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M20 16V6h-5.6L14 4H7.2l12 12h.8"/>`),
		g.Group(children),
	)
}