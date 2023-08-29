package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432.9 0H107.1C94.3 0 83.8 10.4 83.8 23.3V512L270 325.8L456.2 512V23.3c0-12.9-10.4-23.3-23.3-23.3zm-46.5 186.2h-93.1v93.1h-46.5v-93.1h-93.1v-46.5h93.1V46.5h46.5v93.1h93.1v46.6z"/>`),
		g.Group(children),
	)
}