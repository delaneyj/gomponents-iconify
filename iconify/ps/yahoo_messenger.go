package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YahooMessenger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 462q95 0 162.5-67.5T462 232T394.5 69.5T232 2T69.5 69.5T2 232t67.5 162.5T232 462zm59-356q12 0 21 8.5t9 21.5t-9 21.5t-21 8.5q-13 0-21.5-8.5T261 136t8.5-21.5T291 106zm-114 0q12 0 21 8.5t9 21.5t-9 21.5t-21 8.5q-13 0-21.5-8.5T147 136t8.5-21.5T177 106zM70 202h324v60h-1q-5 62-51 105.5T232 411t-110-43.5T71 262h-1v-60zm302 54H88v-37h284v37z"/>`),
		g.Group(children),
	)
}