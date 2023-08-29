package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.9 0H93.1C80.2 0 69.8 10.4 69.8 23.3V512L256 325.8L442.2 512V23.3c0-12.9-10.4-23.3-23.3-23.3zm-46.5 186.2H139.6v-46.5h232.7v46.5z"/>`),
		g.Group(children),
	)
}