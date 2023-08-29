package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faucet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14v2h-3v-1.72l3-.28m0-1c0-1.1-1-2-2.2-2H10v-1H5v11h5v-7.09l9-.91M5 9h5V7l5.36-1.79a.932.932 0 1 0-.61-1.76L5 7v2Z"/>`),
		g.Group(children),
	)
}